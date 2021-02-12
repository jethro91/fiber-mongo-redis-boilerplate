//nolint
package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var CORS_OPTIONS = cors.Config{
	AllowOrigins:     "*",
	AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
	AllowHeaders:     "Content-Type, Content-Type, Credentials, Origin, Accept, Authorization",
	AllowCredentials: true,
}
