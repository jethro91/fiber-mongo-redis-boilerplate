package redisDB

import (
	"context"
	"time"
)

func RemoveSession(userId string) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	err := Client.Del(ctx, "sess"+userId).Err()
	if err != nil {
		return err
	}

	return nil
}
