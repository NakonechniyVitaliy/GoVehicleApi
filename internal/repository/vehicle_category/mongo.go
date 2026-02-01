package vehicle_category

import (
	"context"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	mongoStorage "github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db                *mongo.Database
	vehicleCategories *mongo.Collection
}

func NewMongo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:                db,
		vehicleCategories: db.Collection("vehicle_categories"),
	}
}

func (mng *MongoRepository) GetAll(ctx context.Context) ([]models.VehicleCategory, error) {
	const op = "storage.vehicleCategory.GetAllVehicleCategory"

	result, err := mng.vehicleCategories.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var vehicleCategories []models.VehicleCategory
	if err := result.All(ctx, &vehicleCategories); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return vehicleCategories, nil

}

func (mng *MongoRepository) InsertOrUpdate(ctx context.Context, vehicleCategory models.VehicleCategory) error {
	const op = "storage.vehicleCategory.InsertOrUpdate"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "vehicle_categories")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	vehicleCategory.ID = id

	filter := bson.M{"value": vehicleCategory.Value}

	update := bson.M{
		"$set": vehicleCategory,
	}

	opts := options.Update().SetUpsert(true)

	_, err = mng.vehicleCategories.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
