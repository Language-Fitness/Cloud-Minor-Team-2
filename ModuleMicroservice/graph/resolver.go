package graph

import (
	"Module/graph/model"
	"Module/internal/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver() *Resolver {
	moduleRepository := repository.NewModuleRepository()
	return &Resolver{
		Modules:    []*model.Module{},
		Repository: moduleRepository,
	}
}

type Resolver struct {
	Modules    []*model.Module
	Repository *repository.ModuleRepository
}
