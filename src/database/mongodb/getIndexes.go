package mongoDB

import (
	"context"
	"errors"
	"time"

	"github.com/jethro91/fiber-mongo-redis/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetIndexes(
	databaseId string,
	collectionId string,
) ([]primitive.M, error) {
	results := []bson.M{}

	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return results, errors.New("Cannot getIndexes, collectionId Not Found")
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collectionIndex := Client.Database(databaseId).Collection(collectionId).Indexes()

	cursor, err := collectionIndex.List(ctx)
	if err != nil {
		return results, errors.New("Cannot getIndexes, collectionId Not Found")
	}

	err = cursor.All(context.Background(), &results)
	if err != nil {
		return results, errors.New("Cannot getIndexes, collectionId Not Found")
	}

	return results, nil
}
