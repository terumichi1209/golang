package validator

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// Validator test
type Validator struct {
	validator *validator.Validate
}

// Validate test
func (v *Validator) Validate(i interface{}) error {
	errors := v.validator.Struct(i)
	if errors != nil {
		for _, err := range errors.(validator.ValidationErrors) {
			switch err.ActualTag() {
			case "required":
				return fmt.Errorf("%s は必須です。", err.Field())
			case "email":
				return fmt.Errorf("%s は無効なメールアドレスの形式です。", err.Field())
			}
		}
	}

	return errors
}

// Init test
func Init() *Validator {
	return &Validator{validator: validator.New()}
}
