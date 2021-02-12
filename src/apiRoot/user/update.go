package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis/src/model"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

type updateSchema struct {
	IsVerified bool   `bson:"isVerified" json:"isVerified" validate:"required"`
	Name       string `bson:"name" json:"name" validate:"min=3,max=128,required"`
	Password   string `bson:"password" json:"password" validate:"min=3,max=128,required"`
}

func update(c *fiber.Ctx) error {
	formData := new(updateSchema)
	user := new(model.User)
	var isDocExist bool
	filter := bson.M{
		"_id": c.Params("userId"),
	}

	err := util.ParseBodyAndValidate(c, formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	err = <-user.Exists(c, filter, &isDocExist)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if isDocExist == false {
		errMsg := fmt.Sprintf("User with ID '%s' Not Found", filter["_id"])
		return util.HttpBadRequest(c, errMsg)
	}

	updateDoc := makeUpdateDoc(formData)
	err = <-user.UpdateOne(filter, updateDoc, false)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}

func makeUpdateDoc(formData *updateSchema) *updateSchema {
	updateDoc := formData
	return updateDoc
}
