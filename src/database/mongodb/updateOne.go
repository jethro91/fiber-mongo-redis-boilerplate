package mongoDB

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateOne(
	databaseId string,
	collectionId string,
	filter primitive.M,
	updateDoc primitive.M,
	myUser primitive.M,
	disableTimeStamp bool,
) error {
	if databaseId == "" {
		databaseId = config.MONGO_DATABASE
	}
	if collectionId == "" {
		return errors.New("Cannot updateOne, collectionId Not Found")
	}
	if filter == nil {
		filter = bson.M{}
	}
	if reflect.DeepEqual(filter, bson.M{}) {
		return errors.New("Cannot updateOne, filter is empty")
	}
	if updateDoc == nil {
		updateDoc = bson.M{}
	}
	if reflect.DeepEqual(updateDoc, bson.M{}) {
		return errors.New("Cannot updateOne, updateDoc is empty")
	}
	if myUser == nil {
		myUser = bson.M{}
	}

	if disableTimeStamp == false {
		// Unix Epoch Miliseconds
		updateDoc["updatedAt"] = util.TimeNowUnixEpoch()
		updateDoc["updatedById"] = "system"
		updateDoc["updatedByName"] = "system"
		if myUser["_id"] != nil && myUser["name"] != nil {
			if myUser["_id"] != "" && myUser["name"] != "" {
				updateDoc["updatedById"] = myUser["_id"]
				updateDoc["updatedByName"] = myUser["name"]
			}
		}
	}

	setDoc := bson.M{}
	setDoc["$set"] = updateDoc

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	collection := Client.Database(databaseId).Collection(collectionId)
	_, err := collection.UpdateOne(ctx, filter, setDoc)
	if err != nil {
		return err
	}

	return nil
}
