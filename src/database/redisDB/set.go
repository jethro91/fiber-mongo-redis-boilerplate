package redisDB

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Set(key string, data primitive.M, redisTTL int) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = Client.Set(ctx, key, jsonData, time.Duration(redisTTL)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}
