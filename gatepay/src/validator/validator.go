package validator

import "github.com/go-playground/validator/v10"

var (
	objectValidator = validator.New(validator.WithRequiredStructEnabled())
)

func ValidateObject(i interface{}) error {
	return objectValidator.Struct(i)
}
