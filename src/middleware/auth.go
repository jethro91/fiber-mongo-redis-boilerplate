package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	sessionStore "github.com/jethro91/fiber-mongo-redis/src/model/sessionStoreJWT"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

func IsGuest(c *fiber.Ctx) error {
	if isClientSessionIdExists() == true {
		// return util.HttpBadRequest(c, "You already logged In")
		return c.Next()
	}
	return c.Next()
}

func IsAuth(c *fiber.Ctx) error {
	if isClientSessionIdExists() == false {
		return util.HttpUnauthorized(c, "You must be logged in")
	}
	if isUserIdExists() == false {
		return util.HttpUnauthorized(c, "User Not Found")
	}
	return c.Next()
}

func isClientSessionIdExists() bool {
	if sessionStore.Data["clientSessionId"] == nil {
		return false
	}
	clientSessionId := fmt.Sprintf("%v", sessionStore.Data["clientSessionId"])
	return clientSessionId != ""
}
func isUserIdExists() bool {
	if sessionStore.User["_id"] == nil {
		return false
	}
	userId := fmt.Sprintf("%v", sessionStore.User["_id"])
	return userId != ""
}
