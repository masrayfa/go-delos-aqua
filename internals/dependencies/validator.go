package dependencies

import "github.com/go-playground/validator"

type Validator struct {
	Validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		Validate: validator.New(),
	}
}

func (v *Validator) ValidateStruct(s interface{}) error {
	return v.Validate.Struct(s)
}