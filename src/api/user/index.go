package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis/src/middleware"
)

func RouterGroup(api fiber.Router) {
	home := api.Group("/user")

	home.Get("/my-user", middleware.IsAuth, myUser)
}
