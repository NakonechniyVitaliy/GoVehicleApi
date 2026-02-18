package body_style

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
	db         *mongo.Database
	bodyStyles *mongo.Collection
}

func NewMongoBodyStyleRepo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:         db,
		bodyStyles: db.Collection("body_styles"),
	}
}

func (mng *MongoRepository) Create(ctx context.Context, bodyStyle models.BodyStyle) error {
	const op = "storage.bodyStyle.Create"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "body_styles")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	bodyStyle.ID = id

	_, err = mng.bodyStyles.InsertOne(ctx, bodyStyle)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (mng *MongoRepository) Update(ctx context.Context, bodyStyle models.BodyStyle) error {
	const op = "storage.bodyStyle.UpdateBodyStyle"

	filter := bson.M{
		"id": bodyStyle.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"name":  bodyStyle.Name,
			"value": bodyStyle.Value,
		},
	}

	res, err := mng.bodyStyles.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.MatchedCount == 0 {
		return storage.ErrBodyStyleNotFound
	}

	return nil
}

func (mng *MongoRepository) GetByID(ctx context.Context, bodyStyleID uint16) (*models.BodyStyle, error) {
	const op = "storage.bodyStyle.GetByID"

	filter := bson.D{{"id", bodyStyleID}}

	var bodyStyle models.BodyStyle
	err := mng.bodyStyles.FindOne(ctx, filter).Decode(&bodyStyle)

	switch {
	case err == nil:
		return &bodyStyle, nil
	case err == mongo.ErrNoDocuments:
		return nil, storage.ErrBodyStyleNotFound

	default:
		return nil, fmt.Errorf("%s: %w", op, err)
	}
}

func (mng *MongoRepository) GetAll(ctx context.Context) ([]models.BodyStyle, error) {
	const op = "storage.bodyStyle.UpdateBodyStyle"

	result, err := mng.bodyStyles.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var bodyStyles []models.BodyStyle
	if err := result.All(ctx, &bodyStyles); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return bodyStyles, nil

}

func (mng *MongoRepository) Delete(ctx context.Context, bodyStyleID uint16) error {
	const op = "storage.bodyStyle.Delete"

	filter := bson.D{{"id", bodyStyleID}}

	res, err := mng.bodyStyles.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.DeletedCount == 0 {
		return storage.ErrBodyStyleNotFound
	}
	return nil

}

func (mng *MongoRepository) InsertOrUpdate(ctx context.Context, bodyStyle models.BodyStyle) error {
	const op = "storage.bodyStyle.InsertOrUpdate"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "body_styles")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	bodyStyle.ID = id

	filter := bson.M{"value": bodyStyle.Value}

	update := bson.M{
		"$set": bodyStyle,
	}

	opts := options.Update().SetUpsert(true)

	_, err = mng.bodyStyles.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
