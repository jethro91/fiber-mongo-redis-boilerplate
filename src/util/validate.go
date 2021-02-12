package util

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func ParseBodyAndValidate(c *fiber.Ctx, formData interface{}) error {
	err := c.BodyParser(formData)
	if err != nil {
		return err
	}
	err = Validate(formData)
	if err != nil {
		return err
	}
	return nil
}

func Validate(schema interface{}) error {
	validate := validator.New()
	err := validate.Struct(schema)
	if err != nil {
		fieldError := err.(validator.ValidationErrors)[0]
		const fieldErrMsg = "ValidationError: '%s' Field must be '%s' %s"
		message := fmt.Sprintf(fieldErrMsg, fieldError.Field(), fieldError.Tag(), fieldError.Param())
		return errors.New(message)
	}

	return nil
}
