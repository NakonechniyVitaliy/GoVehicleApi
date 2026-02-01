package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetNextID(ctx context.Context, db *mongo.Database, name string) (int, error) {
	type Counter struct {
		ID  string `bson:"_id"`
		Seq int    `bson:"seq"`
	}

	filter := bson.M{"_id": name}
	update := bson.M{"$inc": bson.M{"seq": 1}}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	var counter Counter
	err := db.Collection("counters").FindOneAndUpdate(ctx, filter, update, opts).Decode(&counter)
	if err != nil {
		return 0, err
	}

	return counter.Seq, nil
}
