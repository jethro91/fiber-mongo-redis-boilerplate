package middleware

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreJWT"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
	"go.mongodb.org/mongo-driver/bson"
)

func IsRoleExists(c *fiber.Ctx) error {
	isArchieved := fmt.Sprintf("%v", sessionStore.User["isArchieved"])
	if isArchieved == "true" {
		return util.HttpForbidden(c, "")
	}
	roles := bson.M{}
	err := util.PrimitiveM(sessionStore.User["roles"], &roles)
	if err != nil {
		return util.HttpError(c, err.Error())
	}
	if reflect.DeepEqual(roles, bson.M{}) {
		return util.HttpForbidden(c, "")
	}

	return c.Next()
}

func IsRoleRoot(c *fiber.Ctx) error {
	isArchieved := fmt.Sprintf("%v", sessionStore.User["isArchieved"])
	if isArchieved == "true" {
		return util.HttpForbidden(c, "")
	}
	roles := bson.M{}
	err := util.PrimitiveM(sessionStore.User["roles"], &roles)
	if err != nil {
		return util.HttpError(c, err.Error())
	}
	isSelectedRoleExists := fmt.Sprintf("%v", roles["root"])
	if isSelectedRoleExists == "true" {
		return util.HttpForbidden(c, "")
	}

	return c.Next()
}
