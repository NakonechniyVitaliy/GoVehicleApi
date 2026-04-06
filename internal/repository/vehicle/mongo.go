package vehicle

import (
	"context"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle/filter"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/_errors"
	mongoStorage "github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db       *mongo.Database
	vehicles *mongo.Collection
}

func NewMongoVehicleRepo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:       db,
		vehicles: db.Collection("vehicles"),
	}
}

func (mng *MongoRepository) Create(ctx context.Context, vehicle models.Vehicle) (*models.Vehicle, error) {
	const op = "storage.vehicle.create"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "vehicles")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	vehicle.ID = id

	res, err := mng.vehicles.InsertOne(ctx, vehicle)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.M{
		"_id": res.InsertedID,
	}
	var createdVehicle models.Vehicle
	err = mng.vehicles.FindOne(ctx, filter).Decode(&createdVehicle)

	return &createdVehicle, nil
}

func (mng *MongoRepository) Update(ctx context.Context, vehicle models.Vehicle, vehicleID uint16) (*models.Vehicle, error) {
	const op = "storage.vehicle.update"

	filter := bson.M{
		"id": vehicleID,
	}

	update := bson.M{
		"$set": bson.M{
			"brand":       vehicle.Brand,
			"driver_type": vehicle.DriverType,
			"gearbox":     vehicle.Gearbox,
			"body_style":  vehicle.BodyStyle,
			"category":    vehicle.Category,
			"mileage":     vehicle.Mileage,
			"model":       vehicle.Model,
			"price":       vehicle.Price,
		},
	}

	res, err := mng.vehicles.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if res.MatchedCount == 0 {
		return nil, _errors.ErrVehicleNotFound
	}

	filter = bson.M{
		"id": vehicleID,
	}
	var updatedVehicle models.Vehicle
	err = mng.vehicles.FindOne(ctx, filter).Decode(&updatedVehicle)

	return &updatedVehicle, nil
}

func (mng *MongoRepository) GetByID(ctx context.Context, vehicleID uint16) (*models.Vehicle, error) {
	const op = "storage.vehicle.get_by_id"

	filter := bson.D{{"id", vehicleID}}

	var vehicle models.Vehicle
	err := mng.vehicles.FindOne(ctx, filter).Decode(&vehicle)

	switch {
	case err == nil:
		return &vehicle, nil
	case err == mongo.ErrNoDocuments:
		return nil, _errors.ErrVehicleNotFound

	default:
		return nil, fmt.Errorf("%s: %w", op, err)
	}
}

func (mng *MongoRepository) GetAll(ctx context.Context) ([]models.Vehicle, error) {
	const op = "storage.vehicle.get_all"

	result, err := mng.vehicles.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var vehicles []models.Vehicle
	if err := result.All(ctx, &vehicles); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return vehicles, nil

}

func (mng *MongoRepository) GetByPage(ctx context.Context, f filter.Filter) ([]models.Vehicle, error) {
	const op = "storage.vehicle.get_by_page"

	skip := int64((f.Page - 1) * f.Limit)
	limit := int64(f.Limit)

	opts := mongoOptions.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := mng.vehicles.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var vehicles []models.Vehicle
	if err := cursor.All(ctx, &vehicles); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return vehicles, nil
}

func (mng *MongoRepository) Delete(ctx context.Context, vehicleID uint16) error {
	const op = "storage.vehicle.delete"

	filter := bson.D{{"id", vehicleID}}

	res, err := mng.vehicles.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.DeletedCount == 0 {
		return _errors.ErrVehicleNotFound
	}
	return nil

}
