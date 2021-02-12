package userAccess

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis/src/middleware"
)

func RouterGroup(apiRoot fiber.Router) {
	userAccess := apiRoot.Group("/user-access")

	userAccess.Get("/list", middleware.IsAuth, list)

	userAccess.Post("/update/:userId", middleware.IsAuth, update)

}
