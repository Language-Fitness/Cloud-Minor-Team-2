package validation

import (
	"Module/graph/model"
	"errors"
)

type Validator struct {
	// Define any fields or rules you need for validation.
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateCreateModuleInput(input *model.ModuleInput) error {
	if input.Name == "" {
		return errors.New("name is required")
	}
	// Implement other validation rules for create here.
	return nil
}

func (v *Validator) ValidateUpdateModuleInput(input *model.ModuleInput) error {
	if input.Name == "" {
		return errors.New("name is required")
	}
	// Implement other validation rules for update here.
	return nil
}
