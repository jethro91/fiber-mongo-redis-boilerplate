package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
)

func CreateCors() func(*fiber.Ctx) error {
	return cors.New(config.CORS_OPTIONS)
}
