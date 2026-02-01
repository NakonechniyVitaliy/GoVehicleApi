package vehicleType

import (
	"context"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	mongoStorage "github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db           *mongo.Database
	vehicleTypes *mongo.Collection
}

func NewMongo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:           db,
		vehicleTypes: db.Collection("vehicle_types"),
	}
}

func (mng *MongoRepository) Create(ctx context.Context, vehicleType models.VehicleType) error {
	const op = "storage.vehicleType.Create"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "vehicle_types")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	vehicleType.ID = id

	_, err = mng.vehicleTypes.InsertOne(ctx, vehicleType)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (mng *MongoRepository) Update(ctx context.Context, vehicleType models.VehicleType) error {
	const op = "storage.vehicleType.UpdateVehicleType"

	filter := bson.M{
		"id": vehicleType.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"ablative":    vehicleType.Ablative,
			"category_id": vehicleType.CategoryID,
			"name":        vehicleType.Name,
			"plural":      vehicleType.Plural,
			"rewrite":     vehicleType.Rewrite,
			"singular":    vehicleType.Singular,
		},
	}

	res, err := mng.vehicleTypes.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.MatchedCount == 0 {
		return storage.ErrVehicleTypeNotFound
	}

	return nil
}

func (mng *MongoRepository) GetByID(ctx context.Context, vehicleTypeID int) (*models.VehicleType, error) {
	const op = "storage.vehicleType.GetByID"

	filter := bson.D{{"id", vehicleTypeID}}

	var vehicleType models.VehicleType
	err := mng.vehicleTypes.FindOne(ctx, filter).Decode(&vehicleType)

	switch {
	case err == nil:
		return &vehicleType, nil
	case err == mongo.ErrNoDocuments:
		return nil, storage.ErrVehicleTypeNotFound

	default:
		return nil, fmt.Errorf("%s: %w", op, err)
	}
}

func (mng *MongoRepository) GetAll(ctx context.Context) ([]models.VehicleType, error) {
	const op = "storage.vehicleType.UpdateVehicleType"

	result, err := mng.vehicleTypes.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var vehicleTypes []models.VehicleType
	if err := result.All(ctx, &vehicleTypes); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return vehicleTypes, nil

}

func (mng *MongoRepository) Delete(ctx context.Context, vehicleTypeID int) error {
	const op = "storage.vehicleType.Delete"

	filter := bson.D{{"id", vehicleTypeID}}

	res, err := mng.vehicleTypes.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.DeletedCount == 0 {
		return storage.ErrVehicleTypeNotFound
	}
	return nil

}

func (mng *MongoRepository) InsertOrUpdate(ctx context.Context, vehicleType models.VehicleType) error {
	const op = "storage.vehicleType.InsertOrUpdate"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "vehicle_types")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	vehicleType.ID = id

	filter := bson.M{"category_id": vehicleType.CategoryID}

	update := bson.M{
		"$set": vehicleType,
	}

	opts := options.Update().SetUpsert(true)

	_, err = mng.vehicleTypes.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
