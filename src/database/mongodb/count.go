package mongoDB

import (
	"context"
	"errors"
	"time"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Count(
	databaseId string,
	collectionId string,
	filter primitive.M,
) (int64, error) {
	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return 0, errors.New("Cannot count, collectionId Not Found")
	}
	if filter == nil {
		filter = bson.M{}
	}

	option := options.Count()

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collection := Client.Database(databaseId).Collection(collectionId)
	count, err := collection.CountDocuments(ctx, filter, option)

	if err != nil {
		return 0, err
	}
	return count, nil
}
