package home

import (
	"github.com/gofiber/fiber/v2"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreJWT"
)

func homeCtrl(c *fiber.Ctx) error {
	return c.Status(200).JSON(sessionStore.User)
}
