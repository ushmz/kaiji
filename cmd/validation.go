package cmd

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func validateParams(p interface{}) error {
	if err := validate.Struct(p); err != nil {
		var errorText []string
		for _, err := range err.(validator.ValidationErrors) {
			errorText = append(errorText, translateValidationError(err))
		}
		return fmt.Errorf("Invalid parameter: %s", strings.Join(errorText, "\n"))
	}
	return nil
}

func translateValidationError(e validator.FieldError) string {
	f := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", f)
	case "max":
		return fmt.Sprintf("%s must be lower than %s", f, e.Param())
	case "min":
		return fmt.Sprintf("%s must be greater then %s", f, e.Param())
	}
	return fmt.Sprintf("%s is not valid %s", e.Field(), e.Value())
}
