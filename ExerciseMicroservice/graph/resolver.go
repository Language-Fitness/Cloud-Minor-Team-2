package graph

import "ExerciseMicroservice/graph/model"
import "ExerciseMicroservice/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.IExerciseService
	Classes []*model.Exercise
}

func NewResolver() *Resolver {
	return &Resolver{
		Service: service.NewExerciseService(),
		Classes: []*model.Exercise{},
	}
}
