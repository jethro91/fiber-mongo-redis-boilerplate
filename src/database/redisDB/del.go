package redisDB

import (
	"context"
	"time"
)

func Del(key string) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	err := Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
