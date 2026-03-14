package user

import (
	"context"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/_errors"
	mongoStorage "github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
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

func (mng *MongoRepository) Create(ctx context.Context, user models.User) error {
	const op = "storage.user.create"

	id, err := mongoStorage.GetNextID(ctx, mng.db, "users")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	user.ID = id

	res, err := mng.users.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.M{
		"_id": res.InsertedID,
	}
	var createdUser models.User
	err = mng.users.FindOne(ctx, filter).Decode(&createdUser)

	return nil
}

func (mng *MongoRepository) GetByID(ctx context.Context, userID uint16) error {
	const op = "storage.user.get_by_id"

	filter := bson.D{{"id", userID}}

	var user models.User
	err := mng.users.FindOne(ctx, filter).Decode(&user)

	switch {
	case err == nil:
		return nil
	case err == mongo.ErrNoDocuments:
		return _errors.ErrUserNotFound

	default:
		return fmt.Errorf("%s: %w", op, err)
	}
}
