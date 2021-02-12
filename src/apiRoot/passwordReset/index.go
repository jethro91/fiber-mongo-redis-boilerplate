package passwordReset

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/middleware"
)

func RouterGroup(apiRoot fiber.Router) {
	passwordReset := apiRoot.Group("/password-reset")

	passwordReset.Get("/list", middleware.IsAuth, list)
	passwordReset.Post("/remove-all", middleware.IsAuth, removeAll)
}
