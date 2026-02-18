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
				"eng":         bson.M{"bsonType": "string"},
				"marka_id":    bson.M{"bsonType": "int"},
				"name":        bson.M{"bsonType": "string"},
				"slang":       bson.M{"bsonType": "string"},
				"value":       bson.M{"bsonType": "int"},
			},
		},
	}

	err := database.CreateCollection(ctx, "brands", options.CreateCollection().SetValidator(brandValidator))
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("create brands collection: %w", err)
	}

	vehicleValidator := bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"id", "brand", "driver_type", "gearbox", "body_style", "category", "mileage", "model", "price"},
			"properties": bson.M{
				"id":          bson.M{"bsonType": "int"},
				"brand":       bson.M{"bsonType": "int"},
				"driver_type": bson.M{"bsonType": "int"},
				"gearbox":     bson.M{"bsonType": "int"},
				"body_style":  bson.M{"bsonType": "string"},
				"category":    bson.M{"bsonType": "int"},
				"mileage":     bson.M{"bsonType": "int"},
				"model":       bson.M{"bsonType": "string"},
				"price":       bson.M{"bsonType": "int"},
			},
		},
	}

	err = database.CreateCollection(ctx, "vehicles", options.CreateCollection().SetValidator(vehicleValidator))
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("create vehicles collection: %w", err)
	}

	vehicleTypes := database.Collection("body_styles")
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
