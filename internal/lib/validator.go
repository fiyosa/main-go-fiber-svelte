package lib

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func NewValidator() *validator.Validate {
	validate = validator.New()
	return validate
}

func GetValidator() *validator.Validate {
	return validate
}

func ValidateRequest(s any) error {
	return validate.Struct(s)
}
