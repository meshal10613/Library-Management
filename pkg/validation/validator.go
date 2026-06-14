package validation

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(v *validator.Validate) *CustomValidator {
	return &CustomValidator{validator: v}
}

// RegisterValidation registers a custom validation function by tag name.
func (cv *CustomValidator) RegisterValidation(tag string, fn validator.Func) error {
	return cv.validator.RegisterValidation(tag, fn)
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}