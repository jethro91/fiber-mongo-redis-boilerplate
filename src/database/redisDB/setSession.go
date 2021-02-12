package redisDB

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSession clientSessionId = userId + : + shortId
func SetSession(clientSessionId string, data primitive.M) error {
	if clientSessionId == "" {
		return errors.New("clientSessionId not found")
	}

	if !strings.Contains(clientSessionId, ":") {
		return errors.New("clientSessionId not valid")
	}
	split := strings.Split(clientSessionId, ":")
	userId := split[0]
	shortId := split[1]

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = Client.HSet(ctx, "sess"+userId, shortId, jsonData).Err()
	if err != nil {
		return err
	}

	return nil
}
