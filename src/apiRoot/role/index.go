package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis/src/middleware"
)

func RouterGroup(apiRoot fiber.Router) {
	role := apiRoot.Group("/role")

	role.Get("/list", middleware.IsAuth, list)
	role.Get("/detail/:roleId", middleware.IsAuth, detail)

	role.Post("/create", middleware.IsAuth, create)
	role.Post("/update/:roleId", middleware.IsAuth, update)
	role.Post("/delete/:roleId", middleware.IsAuth, remove)
}
