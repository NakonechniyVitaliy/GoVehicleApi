package mongo

import (
	"context"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/migrator"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func New(ctx context.Context, cfg *config.Config) (*MongoStorage, error) {
	const op = "storage.mongo.New"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return nil, err
	}

	database := client.Database("core")

	if err := migrator.Run(cfg); err != nil {
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
