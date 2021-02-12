package redisDB

import (
	"context"
	"time"
)

func HDel(groupId string, key string) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	err := Client.HDel(ctx, groupId, key).Err()
	if err != nil {
		return err
	}

	return nil
}
