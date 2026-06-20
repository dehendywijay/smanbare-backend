package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator interface {
	Validate(s interface{}) error
}

type customValidator struct {
	validate *validator.Validate
}

func NewCustomValidator() CustomValidator {
	return &customValidator{
		validate: validator.New(),
	}
}

func (cv *customValidator) Validate(s interface{}) error {
	err := cv.validate.Struct(s)
	if err != nil {
		var errorMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMsgs = append(errorMsgs, fmt.Sprintf("Field '%s' tidak valid pada aturan: %s", err.Field(), err.Tag()))
		}
		return fmt.Errorf(strings.Join(errorMsgs, ", "))
	}
	return nil
}