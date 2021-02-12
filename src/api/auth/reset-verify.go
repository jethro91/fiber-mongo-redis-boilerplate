package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/config"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/model"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util"
	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/util/emailTemplate"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type resetPasswordSchema struct {
	Password             string `json:"password" validate:"min=6,max=72,required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"min=6,max=72,required"`
}

func resetVerify(c *fiber.Ctx) error {
	fullOriginalPath := c.OriginalURL()
	requestURL := config.APP_CLIENT_URL + fullOriginalPath
	id := c.Query("id")
	token := c.Query("token")
	expires := c.Query("expires")
	signature := c.Query("signature")
	isValid := util.IsValidVerificationUrl("/api/auth/reset-verify", id, token, expires, signature, requestURL)
	if isValid == false {
		return util.HttpBadRequest(c, "Password Reset link missmatch")
	}

	// Check Reset Token Validity
	passwordReset := new(model.PasswordReset)
	filter := bson.M{
		"_id": id,
	}
	fields := bson.M{
		"_id":    1,
		"email":  1,
		"userId": 1,
	}
	err := <-passwordReset.FindOne(c, filter, fields)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if passwordReset.ID == "" || passwordReset.UserID == "" {
		return util.HttpBadRequest(c, "Invalid Password Reset link")
	}

	// Start Update the Password
	formData := new(resetPasswordSchema)
	user := new(model.User)
	passwordResetList := new(model.PasswordResetList)

	err = util.ParseBodyAndValidate(c, formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	if formData.Password != formData.PasswordConfirmation {
		return util.HttpBadRequest(c, "Password missmatch")
	}

	updateDoc, err := makeUpdatePasswordDoc(formData)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	filterUser := bson.M{
		"_id": passwordReset.UserID,
	}
	filterPasswordReset := bson.M{
		"userId": passwordReset.UserID,
	}
	updateCh := user.UpdateOne(filterUser, updateDoc, true)
	// Delete reset token
	deleteManyCh := passwordResetList.DeleteMany(filterPasswordReset)
	errUpdate, errDeleteMany := <-updateCh, <-deleteManyCh

	if errUpdate != nil {
		return util.HttpBadRequest(c, errUpdate.Error())
	}
	if errDeleteMany != nil {
		return util.HttpBadRequest(c, errDeleteMany.Error())
	}

	// SendMail
	urlRedirect := config.APP_CLIENT_URL
	html, err := emailTemplate.ResetSuccess(urlRedirect)
	if err != nil {
		return util.HttpBadRequest(c, err.Error())
	}
	err = util.SendMail(
		user.Email,
		"Konfirmasi Ganti Password",
		urlRedirect,
		html,
	)

	return c.Status(200).JSON(bson.M{
		"message": "Password anda berhasil diubah, silahkan login kembali menggunakan password baru",
	})
}

type updatePasswordDoc struct {
	Password string `json:"password" validate:"min=6,max=72,required"`
}

func makeUpdatePasswordDoc(formData *resetPasswordSchema) (*updatePasswordDoc, error) {
	updateDoc := new(updatePasswordDoc)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formData.Password), config.BCRYPT_WORK_FACTOR)
	if err != nil {
		return updateDoc, err
	}
	updateDoc.Password = string(hashedPassword)
	return updateDoc, nil
}
