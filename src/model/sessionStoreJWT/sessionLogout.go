package sessionStoreJWT

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func clearSessionUser(c *fiber.Ctx) {
	Data = bson.M{}
	User = bson.M{}
	c.Set("authorization", "")
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = ""
	cookie.Path = "/"
	c.Cookie(cookie)
}

func SessionLogout(c *fiber.Ctx) {
	clearSessionUser(c)
}
