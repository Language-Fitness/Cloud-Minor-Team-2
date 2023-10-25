package graph

import (
	"Module/graph/model"
	"Module/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver() *Resolver {
	return &Resolver{
		Service: service.NewModuleService(),
		Modules: []*model.Module{},
	}
}

type Resolver struct {
	Service service.IModuleService
	Modules []*model.Module
}
