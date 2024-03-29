package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func GetValidator() *validator.Validate {
	return validate
}

func GetValidationErrorMessages(err error) []string {
	validationErrors := err.(validator.ValidationErrors)
	var errorMessages []string

	for _, validationError := range validationErrors {
		switch validationError.Tag() {
		case "required":
			errorMessages = append(errorMessages, validationError.Field()+" is required")
		case "email":
			errorMessages = append(errorMessages, validationError.Field()+" is not a valid email")
		}
	}

	return errorMessages
}
