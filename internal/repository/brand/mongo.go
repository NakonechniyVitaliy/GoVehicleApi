package brand

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
	db     *mongo.Database
	brands *mongo.Collection
}

func NewMongo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:     db,
		brands: db.Collection("brands"),
	}
}

func (mng *MongoRepository) Create(ctx context.Context, brand models.Brand) error {
	const op = "storage.brand.CreateBrand"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "brands")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	brand.ID = id

	_, err = mng.brands.InsertOne(ctx, brand)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (mng *MongoRepository) Update(ctx context.Context, brand models.Brand) error {
	const op = "storage.brand.UpdateBrand"

	filter := bson.M{
		"id": brand.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"category_id": brand.Category,
			"cnt":         brand.Count,
			"country_id":  brand.Country,
			"eng":         brand.EngName,
			"name":        brand.Name,
			"slang":       brand.Slang,
			"value":       brand.Value,
		},
	}

	res, err := mng.brands.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.MatchedCount == 0 {
		return storage.ErrBrandNotFound
	}

	return nil
}

func (mng *MongoRepository) GetByID(ctx context.Context, brandID int) (*models.Brand, error) {
	const op = "storage.brand.Delete"

	filter := bson.D{{"id", brandID}}

	var brand models.Brand
	err := mng.brands.FindOne(ctx, filter).Decode(&brand)

	switch {
	case err == nil:
		return &brand, nil
	case err == mongo.ErrNoDocuments:
		return nil, storage.ErrBrandNotFound

	default:
		return nil, fmt.Errorf("%s: %w", op, err)
	}
}

func (mng *MongoRepository) GetAll(ctx context.Context) ([]models.Brand, error) {
	const op = "storage.brand.UpdateBrand"

	result, err := mng.brands.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var brands []models.Brand
	if err := result.All(ctx, &brands); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return brands, nil

}

func (mng *MongoRepository) Delete(ctx context.Context, brandID int) error {
	const op = "storage.brand.Delete"

	filter := bson.D{{"id", brandID}}

	res, err := mng.brands.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.DeletedCount == 0 {
		return storage.ErrBrandNotFound
	}
	return nil

}

func (mng *MongoRepository) InsertOrUpdate(ctx context.Context, brand models.Brand) error {
	const op = "storage.brand.InsertOrUpdate"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "brands")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	brand.ID = id

	filter := bson.M{"id": brand.ID}
	update := bson.M{
		"$set": brand,
	}

	opts := options.Update().SetUpsert(true)

	_, err = mng.brands.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
