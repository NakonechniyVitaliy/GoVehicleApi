package mongo

import (
	"context"
	"fmt"
	"strings"

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

	brandValidator := bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"id", "category_id", "cnt", "country_id", "eng", "marka_id", "name", "slang", "value"},
			"properties": bson.M{
				"id":          bson.M{"bsonType": "int"},
				"category_id": bson.M{"bsonType": "int"},
				"cnt":         bson.M{"bsonType": "int"},
				"country_id":  bson.M{"bsonType": "int"},
				"eng":         bson.M{"bsonType": "string", "minLength": 1},
				"marka_id":    bson.M{"bsonType": "int"},
				"name":        bson.M{"bsonType": "string", "minLength": 1},
				"slang":       bson.M{"bsonType": "string", "minLength": 1},
				"value":       bson.M{"bsonType": "int"},
			},
		},
	}

	err := database.CreateCollection(ctx, "brands", options.CreateCollection().SetValidator(brandValidator))
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("create brands collection: %w", err)
	}

	vehicleTypes := database.Collection("vehicle_types")
	_, err = vehicleTypes.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"id": 1},
		Options: options.Index(),
	})
	if err != nil {
		return err
	}

	vehicleCategories := database.Collection("vehicle_categories")
	_, err = vehicleCategories.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"id": 1},
		Options: options.Index(),
	})
	if err != nil {
		return err
	}

	return nil
}
