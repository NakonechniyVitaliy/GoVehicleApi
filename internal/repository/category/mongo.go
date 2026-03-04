package category

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
	categories *mongo.Collection
}

func NewMongoCategoryRepo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:         db,
		categories: db.Collection("categories"),
	}
}

func (mng *MongoRepository) GetByID(ctx context.Context, categoryID uint16) (*models.Category, error) {
	const op = "storage.category.get_by_id"

	filter := bson.D{{"id", categoryID}}

	var category models.Category
	err := mng.categories.FindOne(ctx, filter).Decode(&category)

	switch {
	case err == nil:
		return &category, nil
	case err == mongo.ErrNoDocuments:
		return nil, storage.ErrCategoryNotFound

	default:
		return nil, fmt.Errorf("%s: %w", op, err)
	}
}

func (mng *MongoRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	const op = "storage.category.GetAllCategory"

	result, err := mng.categories.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var categories []models.Category
	if err := result.All(ctx, &categories); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return categories, nil

}

func (mng *MongoRepository) InsertOrUpdate(ctx context.Context, category models.Category) error {
	const op = "storage.category.InsertOrUpdate"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "categories")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	category.ID = id

	filter := bson.M{"value": category.Value}

	update := bson.M{
		"$set": category,
	}

	opts := options.Update().SetUpsert(true)

	_, err = mng.categories.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
