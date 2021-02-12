package sessionStoreRedis

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/database/redisDB"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

func setSessionUser(
	c *fiber.Ctx,
	clientSessionId string,
	newSession primitive.M,
	user primitive.M,
) error {
	Data = newSession
	User = user
	c.Set("authorization", clientSessionId)
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = clientSessionId
	cookie.Path = "/"
	c.Cookie(cookie)
	return redisDB.SetSession(clientSessionId, newSession)
}

func SessionLogin(c *fiber.Ctx, user primitive.M) error {
	now := util.TimeNowUnixEpoch()
	shortIdString, err := shortid.Generate()
	if err != nil {
		return err
	}
	userId := fmt.Sprintf("%v", user["_id"])
	clientSessionId := userId + ":" + shortIdString

	newSession := bson.M{
		"clientSessionId": clientSessionId,
		"user":            user,
		"createdAt":       now,
		"refresh":         now + config.SESSION_REFRESH_TIMEOUT,
		"expires":         now + config.SESSION_EXPIRE_TIMEOUT,
		"redisTTL":        config.REDIS_INIT_TTL,
	}

	err = setSessionUser(c, clientSessionId, newSession, user)
	if err != nil {
		return err
	}
	// Set Session to Redis
	err = redisDB.SetSession(clientSessionId, newSession)
	if err != nil {
		return err
	}

	return nil
}
