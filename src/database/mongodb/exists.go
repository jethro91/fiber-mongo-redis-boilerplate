package mongoDB

import (
	"context"
	"errors"
	"time"

	"github.com/jethro91/fiber-mongo-redis/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Exists(
	databaseId string,
	collectionId string,
	filter primitive.M,
) (bool, error) {

	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return false, errors.New("Cannot check exists, collectionId Not Found")
	}
	if filter == nil {
		filter = bson.M{}
	}

	option := options.Count().SetLimit(1).SetSkip(0)

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collection := Client.Database(databaseId).Collection(collectionId)
	count, err := collection.CountDocuments(ctx, filter, option)

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
