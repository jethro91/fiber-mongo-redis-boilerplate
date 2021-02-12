package mongoDB

import (
	"context"
	"errors"
	"time"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteOne if document didnt exist always return without error
func DeleteMany(
	databaseId string,
	collectionId string,
	filter primitive.M,
) error {
	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return errors.New("Cannot deleteMany, collectionId Not Found")
	}
	if filter == nil {
		filter = bson.M{}
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collection := Client.Database(databaseId).Collection(collectionId)
	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
