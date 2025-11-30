package common

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var _ echo.Validator = &Validator{}

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(value any) error {
	err := v.validator.Struct(value)
	if err != nil {
		return fmt.Errorf("failed validating: %w", err)
	}

	return nil
}
