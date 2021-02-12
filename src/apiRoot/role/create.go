package role

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

type createSchema struct {
	ID   string `bson:"_id" json:"_id" validate:"min=3,max=128,required"`
	Name string `bson:"name" json:"name" validate:"min=3,max=128,required"`
}

func create(c *fiber.Ctx) error {
	formData := new(createSchema)
	role := new(model.Role)
	var isDocExist bool

	err := util.ParseBodyAndValidate(c, formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	filter := bson.M{
		"_id": formData.ID,
	}

	err = <-role.Exists(c, filter, &isDocExist)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	if isDocExist == true {
		errMsg := fmt.Sprintf("Role with ID '%s' exists", filter["_id"])
		return util.HttpBadRequest(c, errMsg)
	}

	newDoc := makeNewDoc(formData)
	err = <-role.InsertOne(newDoc, false)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}
func makeNewDoc(formData *createSchema) *model.Role {
	newDoc := new(model.Role)
	newDoc.ID = formData.ID
	newDoc.Name = formData.Name
	return newDoc
}
