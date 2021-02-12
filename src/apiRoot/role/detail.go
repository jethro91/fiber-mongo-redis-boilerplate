package role

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

func detail(c *fiber.Ctx) error {
	role := new(model.Role)
	filter := bson.M{
		"_id": c.Params("roleId"),
	}
	fields := bson.M{}

	err := <-role.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(role)

}
