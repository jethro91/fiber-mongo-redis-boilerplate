package passwordReset

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

func list(c *fiber.Ctx) error {
	passwordResetList := new(model.PasswordResetList)
	var totalItem int64
	filter := bson.M{}
	fields := bson.M{}

	findCH := passwordResetList.Find(c, filter, fields)
	countCH := passwordResetList.Count(c, filter, &totalItem)

	errFind, errCount := <-findCH, <-countCH

	if errFind != nil {
		return util.HttpBadRequest(c, errFind.Error())
	}
	if errCount != nil {
		return util.HttpBadRequest(c, errCount.Error())
	}

	return c.Status(200).JSON(bson.M{
		"listData":  passwordResetList,
		"totalItem": totalItem,
	})
}
