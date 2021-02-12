package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis/src/model"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

func remove(c *fiber.Ctx) error {
	user := new(model.User)
	var isDocExist bool
	filter := bson.M{
		"_id": c.Params("userId"),
	}

	err := <-user.Exists(c, filter, &isDocExist)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if isDocExist == false {
		errMsg := fmt.Sprintf("User with ID '%s' already deleted", filter["_id"])
		return util.HttpBadRequest(c, errMsg)
	}

	err = <-user.DeleteOne(filter)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}
