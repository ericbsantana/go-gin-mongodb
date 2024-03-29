package validator

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
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

func ParseAndValidateDTO(c *gin.Context, dto interface{}) ([]string, error) {
	if c.Request.ContentLength == 0 {
		return nil, errors.New("request body cannot be empty")
	}

	if err := c.ShouldBindJSON(dto); err != nil {
		return nil, err
	}

	if err := validate.Struct(dto); err != nil {
		validationErrorMessages := GetValidationErrorMessages(err)

		return validationErrorMessages, nil
	}

	return nil, nil
}
