package userAccess

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreJWT"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

type updateSchema struct {
	IsAdd    bool   `bson:"isAdd" json:"isAdd" validate:"required"`
	RoleId   string `bson:"roleId" json:"roleId" validate:"min=1,max=128,required"`
	RoleName string `bson:"roleName" json:"roleName" validate:"min=1,max=128,required"`
}

func update(c *fiber.Ctx) error {
	formData := new(updateSchema)
	user := new(model.User)
	filter := bson.M{
		"_id": c.Params("userId"),
	}
	fields := bson.M{
		"_id":   1,
		"roles": 1,
	}

	err := util.ParseBodyAndValidate(c, formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	err = <-user.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if user.ID == "" {
		errMsg := fmt.Sprintf("User with ID '%s' Not Found", filter["_id"])
		return util.HttpBadRequest(c, errMsg)
	}

	updateDoc := makeUpdateDoc(formData, user)
	err = <-user.UpdateOne(filter, updateDoc, false)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}

type updateDoc struct {
	Roles       primitive.M `bson:"roles" json:"roles,omitempty"`
	RolesAt     int64       `bson:"rolesAt" json:"rolesAt,omitempty"`
	RolesById   string      `bson:"rolesById" json:"rolesById,omitempty"`
	RolesByName string      `bson:"rolesByName" json:"rolesByName,omitempty"`
}

func makeUpdateDoc(formData *updateSchema, user *model.User) *model.User {
	updateDoc := new(model.User)
	newRoles := user.Roles
	myUser := sessionStore.User
	if formData.IsAdd == true {
		newRoles[formData.RoleId] = true
		updateDoc.Roles = newRoles
	} else {
		newRoles[formData.RoleId] = nil
		updateDoc.Roles = newRoles
	}

	updateDoc.RolesAt = util.TimeNowUnixEpoch()
	updateDoc.RolesById = util.PrimitiveToString(myUser["_id"])
	updateDoc.RolesByName = util.PrimitiveToString(myUser["name"])

	return updateDoc
}
