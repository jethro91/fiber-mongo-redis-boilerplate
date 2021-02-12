package redisDB

import (
	"github.com/go-redis/redis/v8"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
)

var (
	Client *redis.Client
)

func CreateRedisDBConnection() {
	client := redis.NewClient(config.REDIS_OPTIONS)
	Client = client
}
