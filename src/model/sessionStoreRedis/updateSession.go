package sessionStoreRedis

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jethro91/fiber-mongo-redis/src/config"
	"github.com/jethro91/fiber-mongo-redis/src/database/redisDB"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

func UpdateSession(c *fiber.Ctx, clientSessionId string, user primitive.M) error {
	now := util.TimeNowUnixEpoch()
	updatedSesion := bson.M{
		"clientSessionId": clientSessionId,
		"user":            user,
		"createdAt":       now,
		"refresh":         now + config.SESSION_REFRESH_TIMEOUT,
		"expires":         now + config.SESSION_EXPIRE_TIMEOUT,
		"redisTTL":        config.REDIS_INIT_TTL,
	}

	// Set Session to Redis
	err := redisDB.SetSession(clientSessionId, updatedSesion)
	if err != nil {
		return err
	}

	Data = updatedSesion
	User = user

	return nil
}
