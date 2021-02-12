package role

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis/src/model"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

func list(c *fiber.Ctx) error {
	roleList := new(model.RoleList)
	var totalItem int64
	filter := bson.M{}
	fields := bson.M{}

	findCH := roleList.Find(c, filter, fields)
	countCH := roleList.Count(c, filter, &totalItem)

	errFind, errCount := <-findCH, <-countCH

	if errFind != nil {
		return util.HttpBadRequest(c, errFind.Error())
	}
	if errCount != nil {
		return util.HttpBadRequest(c, errCount.Error())
	}

	return c.Status(200).JSON(bson.M{
		"listData":  roleList,
		"totalItem": totalItem,
	})
}
