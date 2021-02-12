package sessionStoreRedis

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis/src/database/redisDB"
	"go.mongodb.org/mongo-driver/bson"
)

func clearSessionUser(c *fiber.Ctx) {
	Data = bson.M{}
	User = bson.M{}
	c.Set("authorization", "")
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = ""
	cookie.Path = "/"
	c.Cookie(cookie)
}

func SessionLogout(c *fiber.Ctx) error {
	if Data["clientSessionId"] == nil {
		clearSessionUser(c)
		return nil
	}
	if Data["clientSessionId"] == "" {
		clearSessionUser(c)
		return nil
	}
	clientSessionId := fmt.Sprintf("%v", Data["clientSessionId"])

	split := strings.Split(clientSessionId, ":")
	userId := split[0]

	if userId == "" {
		clearSessionUser(c)
		return c.Next()
	}

	// Remove Data in Redis
	err := redisDB.RemoveSession(userId)
	if err != nil {
		return err
	}

	clearSessionUser(c)
	return nil
}
