package apiRoot

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jethro91/fiber-mongo-redis/src/apiRoot/passwordReset"
	"github.com/jethro91/fiber-mongo-redis/src/apiRoot/role"
	"github.com/jethro91/fiber-mongo-redis/src/apiRoot/user"
)

func CreateApiRouter(app *fiber.App) {

	apiRoot := app.Group("/api-root")
	role.RouterGroup(apiRoot)
	user.RouterGroup(apiRoot)
	passwordReset.RouterGroup(apiRoot)

	apiRoot.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Api Root Router",
		})
	})
}
