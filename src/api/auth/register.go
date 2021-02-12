package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util/emailTemplate"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type registerSchema struct {
	Email                string `bson:"email" json:"email" validate:"email,min=6,max=254,required"`
	Name                 string `bson:"name" json:"name" validate:"min=3,max=254,required"`
	Password             string `bson:"password" json:"password" validate:"min=6,max=72,required"`
	PasswordConfirmation string `bson:"passwordConfirmation" json:"passwordConfirmation" validate:"min=6,max=72,required"`
}

func register(c *fiber.Ctx) error {
	user := new(model.User)
	formData := new(registerSchema)

	err := util.ParseBodyAndValidate(c, formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if formData.Password != formData.PasswordConfirmation {
		return util.HttpBadRequest(c, "Password missmatch")
	}

	filter := bson.M{
		"email": formData.Email,
	}
	fields := bson.M{
		"email": 1,
	}
	err = <-user.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if user.Email != "" {
		errMsg := fmt.Sprintf("User with email '%s' already registred. Try using another email or reset your password", filter["email"])
		return util.HttpBadRequest(c, errMsg)
	}

	// Insert Data
	newDoc, err := makeRegisterDoc(formData, user)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	err = <-user.InsertOne(newDoc, false)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	// SendEmail
	urlToken := util.CreateVerificationUrl("/api/auth/email-verify", newDoc.ID, newDoc.Email)

	html, err := emailTemplate.EmailVerify(urlToken)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	err = util.SendMail(
		newDoc.Email,
		"Verifikasi Email",
		urlToken,
		html,
	)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}

func makeRegisterDoc(formData *registerSchema, user *model.User) (*model.User, error) {
	newDoc := new(model.User)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formData.Password), config.BCRYPT_WORK_FACTOR)
	if err != nil {
		return newDoc, err
	}

	newDoc.ID = util.NewObjectID()
	newDoc.IsArchived = false
	newDoc.VerifiedAt = 0
	newDoc.Email = formData.Email
	newDoc.Name = formData.Name
	newDoc.Password = string(hashedPassword)

	return newDoc, nil
}
