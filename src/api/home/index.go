package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/middleware"
)

func RouterGroup(api fiber.Router) {
	home := api.Group("/home")

	home.Get("/", middleware.IsAuth, homeCtrl)
}
