package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/middleware"
)

func RouterGroup(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/login", middleware.IsGuest, login)
	auth.Post("/logout", middleware.IsAuth, logout)
	auth.Post("/register", middleware.IsGuest, register)

	auth.Get("/email-verify", emailVerify)
	auth.Post("/email-resend", emailResend)

	auth.Post("/reset-request", resetRequest)
	auth.Post("/reset-verify", resetVerify)
}
