package mongoDB

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/jethro91/fiber-mongo-redis/src/config"
	"github.com/jethro91/fiber-mongo-redis/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOne(
	databaseId string,
	collectionId string,
	newDoc primitive.M,
	myUser primitive.M,
	disableTimeStamp bool,
) error {
	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return errors.New("Cannot insertOne, collectionId Not Found")
	}
	if newDoc == nil {
		newDoc = bson.M{}
	}
	if reflect.DeepEqual(newDoc, bson.M{}) {
		return errors.New("Cannot insertOne, newDoc is empty")
	}
	if myUser == nil {
		myUser = bson.M{}
	}

	if disableTimeStamp == false {
		// Unix Epoch Miliseconds
		newDoc["createdAt"] = util.TimeNowUnixEpoch()
		newDoc["createdById"] = "system"
		newDoc["createdByName"] = "system"
		newDoc["updatedAt"] = nil
		newDoc["updatedById"] = ""
		newDoc["updatedByName"] = ""
		if myUser["_id"] != nil && myUser["name"] != nil {
			if myUser["_id"] != "" && myUser["name"] != "" {
				newDoc["createdById"] = myUser["_id"]
				newDoc["createdByName"] = myUser["name"]
			}
		}
	}

	if newDoc["_id"] == "" {
		newObjectId := primitive.NewObjectID()
		newDoc["_id"] = newObjectId.Hex()
	}
	if newDoc["_id"] == nil {
		newObjectId := primitive.NewObjectID()
		newDoc["_id"] = newObjectId.Hex()
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collection := Client.Database(databaseId).Collection(collectionId)
	_, err := collection.InsertOne(ctx, newDoc)
	if err != nil {
		return err
	}

	return nil
}
