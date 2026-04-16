package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/_errors"
	mongoStorage "github.com/NakonechniyVitalii/GoVehicleApi/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db    *mongo.Database
	users *mongo.Collection
}

func NewMongoUserRepo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db:    db,
		users: db.Collection("users"),
	}
}

func (mng *MongoRepository) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	const op = "storage.user.get_by_login"

	filter := bson.D{{"login", login}}

	var user models.User
	err := mng.users.FindOne(ctx, filter).Decode(&user)

	switch {
	case err == nil:
		return &user, nil
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, _errors.ErrUserNotFound

	default:
		return nil, fmt.Errorf("%s: Error to return user %w", op, err)
	}
}

func (mng *MongoRepository) Create(ctx context.Context, user models.User) error {
	const op = "storage.user.create"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "users")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	user.ID = id

	_, err = mng.users.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
