package src

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/api"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/apiRoot"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/middleware"
)

func CreateApp() *fiber.App {
	app := fiber.New()
	app.Use(middleware.CreateCors())
	app.Use(middleware.CheckSessionJWT)

	api.CreateApiRouter(app)
	apiRoot.CreateApiRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Golang Mongo Redis it's works")
	})

	app.Use(middleware.NotFound)

	return app
}
