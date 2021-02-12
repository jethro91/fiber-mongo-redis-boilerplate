package mongoDB

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteOne if document didnt exist always return without error
func DeleteOne(
	databaseId string,
	collectionId string,
	filter primitive.M,
) error {
	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return errors.New("Cannot deleteOne, collectionId Not Found")
	}
	if filter == nil {
		filter = bson.M{}
	}
	if reflect.DeepEqual(filter, bson.M{}) {
		return errors.New("Cannot deleteOne, filter is empty")
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collection := Client.Database(databaseId).Collection(collectionId)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
