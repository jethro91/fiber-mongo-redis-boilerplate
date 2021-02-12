package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis/src/middleware"
)

func RouterGroup(apiRoot fiber.Router) {
	user := apiRoot.Group("/user")

	user.Get("/list", middleware.IsAuth, list)
	user.Get("/list-archived", middleware.IsAuth, listArchived)
	user.Get("/detail/:userId", middleware.IsAuth, detail)

	user.Post("/create", middleware.IsAuth, create)
	user.Post("/update/:userId", middleware.IsAuth, update)
	user.Post("/archive/:userId", middleware.IsAuth, archive)
	user.Post("/delete/:userId", middleware.IsAuth, remove)

}
