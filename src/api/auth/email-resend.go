package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis/src/model"
	"github.com/jethro91/fiber-mongo-redis/src/util"
	"github.com/jethro91/fiber-mongo-redis/src/util/emailTemplate"
	"go.mongodb.org/mongo-driver/bson"
)

type emailResendSchema struct {
	Email string `json:"email" validate:"email,min=6,max=254,required"`
}

func emailResend(c *fiber.Ctx) error {
	formData := new(emailResendSchema)
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
		"email":      1,
		"verifiedAt": 1,
	}
	err = <-user.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	if user.VerifiedAt > 0 {
		return util.HttpBadRequest(c, "Email already verified")
	}
	// SendEmail
	urlToken := util.CreateVerificationUrl("/api/auth/email-verify", user.ID, user.Email)

	html, err := emailTemplate.EmailVerify(urlToken)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	err = util.SendMail(
		user.Email,
		"Verifikasi Email",
		urlToken,
		html,
	)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "If your email address needs to be verified, you will receive an email with the activation link",
	})
}
