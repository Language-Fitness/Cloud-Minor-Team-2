package graph

import (
	"Module/graph/model"
	"Module/internal/repository"
	"Module/internal/validation"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver() *Resolver {
	moduleRepository := repository.NewModuleRepository()
	validator := validation.NewRules()

	return &Resolver{
		Modules:    []*model.Module{},
		Repository: moduleRepository,
		Validator:  validator,
	}
}

type Resolver struct {
	Modules    []*model.Module
	Repository *repository.ModuleRepository
	Validator  *validation.Rules
}
