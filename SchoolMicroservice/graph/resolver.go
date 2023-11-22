package graph

import (
	"example/graph/model"
	"example/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.ISchoolService
	Schools []*model.School
}

func NewResolver() *Resolver {
	return &Resolver{
		Service: service.NewSchoolService(),
		Schools: []*model.School{},
	}
}
