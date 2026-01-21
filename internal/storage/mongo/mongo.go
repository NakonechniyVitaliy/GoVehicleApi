package mongo

import (
	"context"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client *mongo.Client
	brands *mongo.Collection
}

func New(ctx context.Context) (*MongoStorage, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return nil, err
	}

	brandCollection := client.Database("core").Collection("brands")
	brandCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"marka_id": 1},
		Options: options.Index().SetUnique(true),
	})

	return &MongoStorage{
		client: client,
		brands: brandCollection,
	}, nil
}

func (mng *MongoStorage) NewBrand(ctx context.Context, brand models.Brand) error {

	_, err := mng.brands.InsertOne(ctx, brand)
	if err != nil {
		return err
	}
	return nil
}

func (mng *MongoStorage) UpdateBrand(ctx context.Context, brand models.Brand) error {
	const op = "storage.brand.UpdateBrand"

	filter := bson.M{
		"marka_id": brand.Marka,
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

func (mng *MongoStorage) GetBrand(ctx context.Context, brandID int) (*models.Brand, error) {
	filter := bson.D{{"marka_id", brandID}}

	var brand models.Brand
	err := mng.brands.FindOne(ctx, filter).Decode(&brand)

	switch {
	case err == nil:
		return &brand, nil
	case err == mongo.ErrNoDocuments:
		return nil, storage.ErrBrandNotFound

	default:
		return nil, err
	}
}

func (mng *MongoStorage) DeleteBrand(ctx context.Context, brandID int) error {
	filter := bson.D{{"marka_id", brandID}}

	res, err := mng.brands.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return storage.ErrBrandNotFound
	}
	return nil

}

func (mng *MongoStorage) RefreshBrands() error {
	return nil
}
