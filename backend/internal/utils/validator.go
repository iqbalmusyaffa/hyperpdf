package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct validates a struct against validator tags
func ValidateStruct(s interface{}) []string {
	var errMessages []string
	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errMessages = append(errMessages, fmt.Sprintf("%s is required", err.Field()))
			case "oneof":
				errMessages = append(errMessages, fmt.Sprintf("%s must be one of: %s", err.Field(), err.Param()))
			case "min":
				errMessages = append(errMessages, fmt.Sprintf("%s must be at least %s", err.Field(), err.Param()))
			case "max":
				errMessages = append(errMessages, fmt.Sprintf("%s must be at most %s", err.Field(), err.Param()))
			default:
				errMessages = append(errMessages, fmt.Sprintf("%s is not valid (%s)", err.Field(), err.Tag()))
			}
		}
	}
	return errMessages
}
