package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	sessionStore "github.com/jethro91/fiber-mongo-redis-boilerplate/src/model/sessionStoreJWT"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
)

type loginSchema struct {
	Email    string `bson:"email" json:"email" validate:"email,min=6,max=254,required"`
	Password string `bson:"password" json:"password" validate:"min=6,max=72,required"`
}

func login(c *fiber.Ctx) error {
	user := new(model.User)
	formData := new(loginSchema)

	err := util.ParseBodyAndValidate(c, formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	filter := bson.M{
		"email": formData.Email,
	}
	fields := bson.M{}
	err = <-user.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if user.ID == "" {
		return util.HttpBadRequest(c, "Incorrect email or user is not registered")
	}

	userBson := bson.M{}
	err = util.PrimitiveM(user, &userBson)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	// Compare Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
	if err != nil {
		return util.HttpBadRequest(c, "Incorrect password")
	}

	err = sessionStore.SessionLogin(c, userBson)
	if err != nil {
		return util.HttpBadRequest(c, "Can't Login")
	}

	return c.Status(200).JSON(bson.M{
		"message":   "OK",
		"sessionId": sessionStore.Data["clientSessionId"],
	})
}
