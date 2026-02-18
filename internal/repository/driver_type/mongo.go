package driver_type

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
	db          *mongo.Database
	driverTypes *mongo.Collection
}

func NewMongoDriverTypeRepo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:          db,
		driverTypes: db.Collection("driver_types"),
	}
}

func (mng *MongoRepository) GetAll(ctx context.Context) ([]models.DriverType, error) {
	const op = "storage.driverType.GetAllDriverType"

	result, err := mng.driverTypes.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var driverTypes []models.DriverType
	if err := result.All(ctx, &driverTypes); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return driverTypes, nil

}

func (mng *MongoRepository) InsertOrUpdate(ctx context.Context, driverType models.DriverType) error {
	const op = "storage.driverType.InsertOrUpdate"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "driver_types")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	driverType.ID = id

	filter := bson.M{"value": driverType.Value}

	update := bson.M{
		"$set": driverType,
	}

	opts := options.Update().SetUpsert(true)

	_, err = mng.driverTypes.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
