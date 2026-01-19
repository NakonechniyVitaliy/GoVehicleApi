package mongo

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client *mongo.Client
}

func New(ctx context.Context) (*MongoStorage, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))

	if err != nil {
		return nil, err
	}

	return &MongoStorage{client: client}, nil
}

func (mng *MongoStorage) NewBrand(brand models.Brand, ctx context.Context) error {
	collection := mng.client.Database("core").Collection("brand")

	_, err := collection.InsertOne(ctx, brand)
	if err != nil {
		return err
	}
	return nil
}

func (mng *MongoStorage) RefreshBrands() error {
	return nil
}
