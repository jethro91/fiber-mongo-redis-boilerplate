package mongoDB

import (
	"context"
	"errors"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jethro91/fiber-mongo-redis/src/config"
)

// FindOne if empty result, return empty bson.M{}
func FindOne(
	databaseId string,
	collectionId string,
	filter primitive.M,
	fields primitive.M,
) (primitive.M, error) {
	var result bson.M
	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return result, errors.New("Cannot findone, collectionId Not Found")
	}
	if filter == nil {
		filter = bson.M{}
	}
	if fields == nil {
		fields = bson.M{}
	}

	option := options.FindOne()
	if !(reflect.DeepEqual(fields, bson.M{})) {
		option = options.FindOne().SetProjection(fields)
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collection := Client.Database(databaseId).Collection(collectionId)
	err := collection.FindOne(ctx, filter, option).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return result, nil
		}
		return result, err
	}

	return result, nil
}
