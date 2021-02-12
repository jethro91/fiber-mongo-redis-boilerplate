//nolint
package config

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var REDIS_HOST = os.Getenv("REDIS_HOST")
var REDIS_PORT = os.Getenv("REDIS_PORT")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

var REDIS_OPTIONS = &redis.Options{
	Addr:     REDIS_HOST + ":" + REDIS_PORT,
	Password: REDIS_PASSWORD,
	DB:       0,
}
