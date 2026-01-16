package mongo

import (
	"context"
	"fmt"
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

func (mng *MongoStorage) NewBrand() string {
	return ""
}

func (mng *MongoStorage) RefreshBrands() string {
	return ""
}
