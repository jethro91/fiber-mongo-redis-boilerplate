package redisDB

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HSet(groupId string, key string, data primitive.M) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = Client.HSet(ctx, groupId, key, jsonData).Err()
	if err != nil {
		return err
	}

	return nil
}
