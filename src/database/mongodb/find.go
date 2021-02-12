package mongoDB

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/jethro91/fiber-mongo-redis/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find(
	databaseId string,
	collectionId string,
	filter primitive.M,
	fields primitive.M,
	limit int64,
	page int64,
	sort string,
	sortDir string,
) ([]primitive.M, error) {
	results := []bson.M{{}}
	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return results, errors.New("Cannot find, collectionId Not Found")
	}
	if filter == nil {
		filter = bson.M{}
	}
	if !(filter["search"] == "" || filter["search"] == nil) {
		filter["$text"] = bson.M{
			"$search": filter["search"],
		}
		filter["search"] = nil
	}
	if fields == nil {
		fields = bson.M{}
	}
	if limit <= 0 {
		limit = 10
	}
	if sort == "" {
		sort = "createdAt"
	}
	if sortDir == "" {
		sortDir = "asc"
	}

	if page <= 0 {
		page = 1
	}

	var skip int64 = 0
	if page > -1 {
		skip = (page - 1) * limit
	} else {
		skip = 0
	}

	sortBson := bson.M{}
	if sortDir == "desc" {
		sortBson[sort] = -1
	} else {
		sortBson[sort] = 1
	}

	option := options.Find().SetSort(sortBson).SetSkip(skip).SetLimit(limit)
	if !(reflect.DeepEqual(fields, bson.M{})) {
		option = options.Find().SetSort(sortBson).SetSkip(skip).SetLimit(limit).SetProjection(fields)
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collection := Client.Database(databaseId).Collection(collectionId)
	cursor, err := collection.Find(ctx, filter, option)

	if err != nil {
		return results, err
	}

	err = cursor.All(context.Background(), &results)
	if err != nil {
		return results, err
	}

	return results, nil
}
