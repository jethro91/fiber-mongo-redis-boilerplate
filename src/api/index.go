package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jethro91/fiber-mongo-redis/src/api/auth"
	"github.com/jethro91/fiber-mongo-redis/src/api/home"
	"github.com/jethro91/fiber-mongo-redis/src/api/user"
)

func CreateApiRouter(app *fiber.App) {

	api := app.Group("/api")
	auth.RouterGroup(api)
	home.RouterGroup(api)
	user.RouterGroup(api)

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Api Router",
		})
	})
}
