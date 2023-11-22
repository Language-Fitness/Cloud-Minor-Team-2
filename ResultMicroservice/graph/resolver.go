package graph

import "ResultMicroservice/internal/service"
import "ResultMicroservice/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.IResultService
	Results []*model.Result
}

func NewResolver() *Resolver {
	return &Resolver{
		Service: service.NewResultService(),
		Results: []*model.Result{},
	}
}
