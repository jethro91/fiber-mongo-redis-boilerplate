package user

import (
	"github.com/gofiber/fiber/v2"
	sessionStore "github.com/jethro91/fiber-mongo-redis/src/model/sessionStoreJWT"
)

func myUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(sessionStore.User)
}
