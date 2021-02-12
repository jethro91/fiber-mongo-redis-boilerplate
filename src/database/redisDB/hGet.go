package redisDB

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HGet(groupId string, key string) (primitive.M, error) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	val, err := Client.HGet(ctx, groupId, key).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return bson.M{}, nil
		}
		return nil, err
	}

	result := bson.M{}
	err = json.Unmarshal([]byte(val), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
