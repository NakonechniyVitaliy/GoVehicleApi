package vehicleType

import (
	"context"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	vehicleTypes *mongo.Collection
}

func NewMongo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		vehicleTypes: db.Collection("vehicle_types"),
	}
}

func (mng *MongoRepository) Create(ctx context.Context, vehicleType models.VehicleType) error {

	_, err := mng.vehicleTypes.InsertOne(ctx, vehicleType)
	if err != nil {
		return err
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
	filter := bson.D{{"id", vehicleTypeID}}

	var vehicleType models.VehicleType
	err := mng.vehicleTypes.FindOne(ctx, filter).Decode(&vehicleType)

	switch {
	case err == nil:
		return &vehicleType, nil
	case err == mongo.ErrNoDocuments:
		return nil, storage.ErrVehicleTypeNotFound

	default:
		return nil, err
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
	filter := bson.D{{"id", vehicleTypeID}}

	res, err := mng.vehicleTypes.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return storage.ErrVehicleTypeNotFound
	}
	return nil

}
