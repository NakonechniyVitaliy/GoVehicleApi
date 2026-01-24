package brand

import "go.mongodb.org/mongo-driver/mongo"

type MongoRepository struct {
	vehicleTypes *mongo.Collection
}

func NewMongo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		vehicleTypes: db.Collection("vehicle_types"),
	}
}

func (mng *MongoRepository) NewVehicleType() error {
	return nil
}

func (mng *MongoRepository) GetVehicleType() error {
	return nil
}

func (mng *MongoRepository) DeleteVehicleType() error {
	return nil
}
