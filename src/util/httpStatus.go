package util

import (
	"github.com/gofiber/fiber/v2"
)

func HttpError(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "HTTP Error"
	}
	return c.Status(500).JSON(fiber.Map{
		"message": message,
	})
}

func HttpBadRequest(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Bad Request"
	}
	return c.Status(400).JSON(fiber.Map{
		"message": message,
	})
}

func HttpUnauthorized(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Unauthorized"
	}
	return c.Status(401).JSON(fiber.Map{
		"message": message,
	})
}

func HttpForbidden(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "You don't have access in this page"
	}
	return c.Status(403).JSON(fiber.Map{
		"message": message,
	})
}

func HttpMaintenance(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Maintenance Server"
	}
	return c.Status(503).JSON(fiber.Map{
		"message": message,
	})
}

func HttpPaymentRequired(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Please make your payment before using this functions"
	}
	return c.Status(402).JSON(fiber.Map{
		"message": message,
	})
}
