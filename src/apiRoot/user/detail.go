package user

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis/src/model"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

func detail(c *fiber.Ctx) error {
	user := new(model.User)
	filter := bson.M{
		"_id": c.Params("userId"),
	}
	fields := bson.M{}

	err := <-user.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(user)

}
