package redisDB

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSession clientSessionId = userId + : + shortId
func GetSession(clientSessionId string) (primitive.M, error) {
	if clientSessionId == "" {
		return nil, errors.New("clientSessionId not found")
	}
	if !strings.Contains(clientSessionId, ":") {
		return nil, errors.New("clientSessionId not valid")
	}
	split := strings.Split(clientSessionId, ":")
	userId := split[0]
	shortId := split[1]

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	val, err := Client.HGet(ctx, "sess"+userId, shortId).Result()

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
