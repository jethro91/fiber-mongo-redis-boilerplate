package mongoDB

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/jethro91/fiber-mongo-redis/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	indexOptions "go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndex(
	databaseId string,
	collectionId string,
	name string,
	fields primitive.M,
	weight primitive.M,
) error {
	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return errors.New("Cannot createIndex, collectionId Not Found")
	}
	if fields == nil {
		fields = bson.M{}
	}
	if reflect.DeepEqual(fields, bson.M{}) {
		return errors.New("Cannot createIndex, fields is empty")
	}
	if weight == nil {
		weight = bson.M{}
	}

	opt := indexOptions.Index()
	if name != "" {
		opt.SetName(name)
	}
	if !reflect.DeepEqual(weight, bson.M{}) {
		opt.SetWeights(weight)
	}

	setIndex := mongo.IndexModel{
		Keys:    fields,
		Options: opt,
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collectionIndex := Client.Database(databaseId).Collection(collectionId).Indexes()

	_, err := collectionIndex.CreateOne(ctx, setIndex)
	if err != nil {
		return err
	}
	return nil
}
