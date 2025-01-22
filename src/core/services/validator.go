package services

import (
	"github.com/go-playground/validator/v10"
)

type ValidatorService struct {
	validate *validator.Validate
}

func NewValidatorService() *ValidatorService {
	return &ValidatorService{
		validate: validator.New(),
	}
}

// ValidateStruct validates a given struct
func (v *ValidatorService) ValidateStruct(obj interface{}) error {
	return v.validate.Struct(obj)
}

// RegisterCustomValidation adds a custom validation rule
func (v *ValidatorService) RegisterCustomValidation(tag string, fn validator.Func) error {
	return v.validate.RegisterValidation(tag, fn)
}