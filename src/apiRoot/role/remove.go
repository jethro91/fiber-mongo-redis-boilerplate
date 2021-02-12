package role

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

func remove(c *fiber.Ctx) error {
	role := new(model.Role)
	var isDocExist bool
	filter := bson.M{
		"_id": c.Params("roleId"),
	}

	err := <-role.Exists(c, filter, &isDocExist)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if isDocExist == false {
		errMsg := fmt.Sprintf("Role with ID '%s' already deleted", filter["_id"])
		return util.HttpBadRequest(c, errMsg)
	}

	err = <-role.DeleteOne(filter)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}
