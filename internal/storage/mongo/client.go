package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	DB     *mongo.Database
	client *mongo.Client
}

func New(ctx context.Context) (*MongoStorage, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return nil, err
	}

	return &MongoStorage{
		DB:     client.Database("core"),
		client: client,
	}, nil
}

func (mng *MongoStorage) CloseDB() error {
	return mng.client.Disconnect(context.Background())
}

func (mng *MongoStorage) GetName() string {
	return "mongo"
}

func (mng *MongoStorage) GetDatabase() *mongo.Database {
	return mng.DB
}

//func New(ctx context.Context) (*MongoStorage, error) {
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
//	if err != nil {
//		return nil, err
//	}
//
//	brandCollection := client.Database("core").Collection("brands")
//	brandCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
//		Keys:    bson.M{"marka_id": 1},
//		Options: options.Index().SetUnique(true),
//	})
//
//	vehicleTypesCollection := client.Database("core").Collection("vehicle_types")
//	vehicleTypesCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
//		Keys:    bson.M{"id": 1},
//		Options: options.Index().SetUnique(true),
//	})
//
//	return &MongoStorage{
//		client:       client,
//		brands:       brandCollection,
//		vehicleTypes: vehicleTypesCollection,
//	}, nil
//}
