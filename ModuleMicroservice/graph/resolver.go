package graph

import (
	"Module/graph/model"
	"Module/internal/auth"
	"Module/internal/database"
	"Module/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.IModuleService
	Modules []*model.Module
	Policy  *auth.Policy
}

func NewResolver() *Resolver {
	collection, _ := database.GetCollection()

	return &Resolver{
		Service: service.NewModuleService(collection),
		Modules: []*model.Module{},
		Policy:  auth.NewPolicy(collection),
	}
}
