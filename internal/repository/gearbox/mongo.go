package gearbox

import (
	"context"
	"errors"
	"fmt"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/_errors"
	mongoStorage "github.com/NakonechniyVitalii/GoVehicleApi/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db        *mongo.Database
	gearboxes *mongo.Collection
}

func NewMongoGearboxRepo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:        db,
		gearboxes: db.Collection("gearboxes"),
	}
}

func (mng *MongoRepository) GetByID(ctx context.Context, gearboxID uint16) (*models.Gearbox, error) {
	const op = "storage.gearbox.get_by_id"

	filter := bson.D{{"id", gearboxID}}

	var gearbox models.Gearbox
	err := mng.gearboxes.FindOne(ctx, filter).Decode(&gearbox)

	switch {
	case err == nil:
		return &gearbox, nil
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, _errors.ErrGearboxNotFound

	default:
		return nil, fmt.Errorf("%s: %w", op, err)
	}
}

func (mng *MongoRepository) GetAll(ctx context.Context) ([]models.Gearbox, error) {
	const op = "storage.gearbox.get_all"

	result, err := mng.gearboxes.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var gearboxes []models.Gearbox
	if err := result.All(ctx, &gearboxes); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return gearboxes, nil

}

func (mng *MongoRepository) InsertOrUpdate(ctx context.Context, gearbox models.Gearbox) error {
	const op = "storage.gearbox.insert_or_update"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "gearboxes")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	gearbox.ID = id

	filter := bson.M{"value": gearbox.Value}

	update := bson.M{
		"$set": gearbox,
	}

	opts := options.Update().SetUpsert(true)

	_, err = mng.gearboxes.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
