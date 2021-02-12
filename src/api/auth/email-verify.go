package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis/src/config"
	"github.com/jethro91/fiber-mongo-redis/src/model"
	"github.com/jethro91/fiber-mongo-redis/src/util"
	"go.mongodb.org/mongo-driver/bson"
)

func emailVerify(c *fiber.Ctx) error {
	fullOriginalPath := c.OriginalURL()
	requestURL := config.APP_CLIENT_URL + fullOriginalPath
	id := c.Query("id")
	token := c.Query("token")
	expires := c.Query("expires")
	signature := c.Query("signature")
	isValid := util.IsValidVerificationUrl("/api/auth/email-verify", id, token, expires, signature, requestURL)
	if isValid == false {
		return util.HttpBadRequest(c, "Activation link missmatch")
	}

	user := new(model.User)
	filter := bson.M{
		"_id": id,
	}
	fields := bson.M{
		"_id":   1,
		"email": 1,
	}
	err := <-user.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if user.ID == "" {
		return util.HttpBadRequest(c, "Invalid activation link")
	}
	if user.VerifiedAt != 0 {
		return util.HttpBadRequest(c, "Email already verified")
	}

	updateDoc := makeVerifyDoc()
	err = <-user.UpdateOne(filter, updateDoc, true)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}

	return c.Status(200).JSON(bson.M{
		"message": "OK",
	})
}

type verifyDoc struct {
	VerifiedAt int64 `bson:"verifiedAt" json:"verifiedAt"`
}

func makeVerifyDoc() *verifyDoc {
	updateDoc := new(verifyDoc)
	updateDoc.VerifiedAt = util.TimeNowUnixEpoch()
	return updateDoc
}
