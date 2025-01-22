package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
	Param       string `json:"param,omitempty"` // Include parameter like "min=3"
}

func ParseValidationErrors(err error) []ValidationError {
	var errors []ValidationError
	for _, fieldErr := range err.(validator.ValidationErrors) {
		validationError := ValidationError{
			FailedField: fieldErr.Field(),
			Tag:         fieldErr.Tag(),
			Value:       fmt.Sprintf("%v", fieldErr.Value()),
		}

		// Extract tag-specific parameters (e.g., "min", "max", "len")
		if param := fieldErr.Param(); param != "" {
			validationError.Param = param
		}

		errors = append(errors, validationError)
	}
	return errors
}
