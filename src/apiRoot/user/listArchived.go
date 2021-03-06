package user

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

func listArchived(c *fiber.Ctx) error {
	userList := new(model.UserList)
	var totalItem int64
	filter := bson.M{
		"isArchived": true,
	}
	fields := bson.M{}

	findCH := userList.Find(c, filter, fields)
	countCH := userList.Count(c, filter, &totalItem)

	errFind, errCount := <-findCH, <-countCH

	if errFind != nil {
		return util.HttpBadRequest(c, errFind.Error())
	}
	if errCount != nil {
		return util.HttpBadRequest(c, errCount.Error())
	}

	return c.Status(200).JSON(bson.M{
		"listData":  userList,
		"totalItem": totalItem,
	})
}
