package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func New(ctx context.Context) (*MongoStorage, error) {
	const op = "storage.mongo.New"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return nil, err
	}

	database := client.Database("core")

	if err := migrate(ctx, database); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &MongoStorage{
		DB:     database,
		Client: client,
	}, nil
}

func (mng *MongoStorage) CloseDB() error {
	return mng.Client.Disconnect(context.Background())
}

func (mng *MongoStorage) GetName() string {
	return "mongo"
}
func migrate(ctx context.Context, database *mongo.Database) error {

	brands := database.Collection("brands")
	_, err := brands.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"marka_id": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	vehicleTypes := database.Collection("vehicle_types")
	_, err = vehicleTypes.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"id": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	return nil
}
