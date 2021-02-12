package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/database/mongoDB"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/database/mongoIndex"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/database/redisDB"
)

func main() {
	app := src.CreateApp()

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	mongoDB.CreateMongoDBConnection()
	defer mongoDB.Client.Disconnect(ctx)

	err := mongoIndex.CreateMongoIndexes()
	if err != nil {
		fmt.Println(err)
		return
	}

	redisDB.CreateRedisDBConnection()

	app.Listen(config.APP_LISTEN)
}
