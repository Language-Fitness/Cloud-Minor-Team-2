package graph

import (
	"Module/graph/model"
	"Module/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.IModuleService
	Modules []*model.Module
}

func NewResolver() *Resolver {
	return &Resolver{
		Service: service.NewModuleService(),
		Modules: []*model.Module{},
	}
}
