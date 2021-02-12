package auth

import (
	"github.com/gofiber/fiber/v2"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreJWT"
	"go.mongodb.org/mongo-driver/bson"
)

func logout(c *fiber.Ctx) error {
	sessionStore.SessionLogout(c)

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}
