package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis/src/model"
	"github.com/jethro91/fiber-mongo-redis/src/util"
)

type createSchema struct {
	IsVerified bool   `bson:"isVerified" json:"isVerified" validate:"required"`
	Email      string `json:"email" validate:"email,min=6,max=254,required"`
	Name       string `bson:"name" json:"name" validate:"min=3,max=128,required"`
	Password   string `bson:"password" json:"password" validate:"min=3,max=128,required"`
}

func create(c *fiber.Ctx) error {
	formData := new(createSchema)
	user := new(model.User)
	var isDocExist bool

	err := util.ParseBodyAndValidate(c, formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	filter := bson.M{
		"email": formData.Email,
	}

	err = <-user.Exists(c, filter, &isDocExist)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if isDocExist == true {
		errMsg := fmt.Sprintf("User with email '%s' already exists", filter["email"])
		return util.HttpBadRequest(c, errMsg)
	}

	newDoc := makeNewDoc(formData)
	err = <-user.InsertOne(newDoc, false)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}
func makeNewDoc(formData *createSchema) *model.User {
	newDoc := new(model.User)
	newDoc.VerifiedAt = 0
	if formData.IsVerified == true {
		newDoc.VerifiedAt = util.TimeNowUnixEpoch()
	}
	newDoc.Email = formData.Email
	newDoc.Password = formData.Password
	return newDoc
}
