package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreJWT"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

type archiveSchema struct {
	IsArchived bool `bson:"isArchived" json:"isArchived" validate:"required"`
}

func archive(c *fiber.Ctx) error {
	formData := new(archiveSchema)
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
		errMsg := fmt.Sprintf("Role with ID '%s' Not Found", filter["_id"])
		return util.HttpBadRequest(c, errMsg)
	}

	updateDoc := makeArchiveDoc(formData)
	err = <-user.UpdateOne(filter, updateDoc, true)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}

type archiveDoc struct {
	IsArchived bool `bson:"isArchived" json:"isArchived"`
}

func makeArchiveDoc(formData *archiveSchema) *model.User {
	updateDoc := new(model.User)
	myUser := sessionStore.User
	if formData.IsArchived == true {
		updateDoc.IsArchived = true
		updateDoc.DeletedAt = util.TimeNowUnixEpoch()
		updateDoc.DeletedById = util.PrimitiveToString(myUser["_id"])
		updateDoc.DeletedByName = util.PrimitiveToString(myUser["name"])
	} else {
		updateDoc.IsArchived = false
		updateDoc.DeletedAt = 0
		updateDoc.DeletedById = ""
		updateDoc.DeletedByName = ""
	}

	return updateDoc
}
