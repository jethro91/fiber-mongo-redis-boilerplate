package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis/src/config"
	"github.com/jethro91/fiber-mongo-redis/src/model"
	"github.com/jethro91/fiber-mongo-redis/src/util"
	"github.com/jethro91/fiber-mongo-redis/src/util/emailTemplate"
	"go.mongodb.org/mongo-driver/bson"
)

type forgotPasswordSchema struct {
	Email string `json:"email" validate:"email,min=6,max=254,required"`
}

func resetRequest(c *fiber.Ctx) error {
	formData := new(forgotPasswordSchema)
	user := new(model.User)

	err := util.ParseBodyAndValidate(c, formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	filter := bson.M{
		"email": formData.Email,
	}
	fields := bson.M{
		"_id":        1,
		"name":       1,
		"email":      1,
		"verifiedAt": 1,
	}

	err = <-user.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	if user.ID == "" {
		return util.HttpBadRequest(c, "User did not exists")
	}

	// Create and Save Token
	newTokenByte, err := util.GenerateRandomBytes()
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	newDoc := makePasswordResetDoc(user, string(newTokenByte))
	err = <-newDoc.InsertOne(newDoc, true)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	// Send Mail
	urlToken := util.CreateVerificationUrl("/api/auth/reset-verify", newDoc.ID, user.Email)
	html, err := emailTemplate.ResetRequest(urlToken)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	err = util.SendMail(
		user.Email,
		"Konfirmasi Ganti Password",
		urlToken,
		html,
	)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "If you have an account with us, you will receive an email with a link to reset your password",
	})
}

func makePasswordResetDoc(user *model.User, newToken string) *model.PasswordReset {
	newDoc := new(model.PasswordReset)
	hasedToken := util.CreateSha256(string(newToken))
	newDoc.ID = util.NewObjectID()
	newDoc.UserID = user.ID
	newDoc.Token = hasedToken
	newDoc.ExpiredAt = util.TimeNowUnixEpoch() + config.PASSWORD_RESET_TIMEOUT

	return newDoc
}
