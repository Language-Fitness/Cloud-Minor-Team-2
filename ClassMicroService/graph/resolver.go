package graph

import (
	"Class/graph/model"
	"Class/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.IClassService
	Classes []*model.Class
}

func NewResolver() *Resolver {
	return &Resolver{
		Service: service.NewClassService(),
		Classes: []*model.Class{},
	}
}
