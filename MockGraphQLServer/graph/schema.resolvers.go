package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"
	"example/graph/model"
	"github.com/google/uuid"
	"time"
)

// CreateResult is the resolver for the createResult field.
func (r *mutationResolver) CreateResult(ctx context.Context, input *model.NewResult) (*model.Result, error) {

	newResult := &model.Result{
		ID:          uuid.New().String(),
		ExerciseID:  input.ExerciseID,
		UserID:      input.UserID,
		ClassID:     input.ClassID,
		ModuleID:    input.ModuleID,
		Input:       input.Input,
		Result:      input.Result,
		CreatedAt:   time.Now().String(),
		SoftDeleted: false,
	}

	r.results = append(r.results, newResult)

	return newResult, nil
}

// GetModuleByID is the resolver for the getModuleById field.
func (r *queryResolver) GetModuleByID(ctx context.Context, id string) (*model.Module, error) {
	if len(r.modules) == 0 {
		r.init()
	}

	for _, obj := range r.modules {
		if obj.ID == id {
			return obj, nil
		}
	}
	return nil, errors.New("no module wound with this id")
}

// GetAllModules is the resolver for the getAllModules field.
func (r *queryResolver) GetAllModules(ctx context.Context) ([]*model.Module, error) {
	if len(r.modules) == 0 {
		r.init()
	}

	return r.modules, nil
}

// GetClassByID is the resolver for the getClassById field.
func (r *queryResolver) GetClassByID(ctx context.Context, id string) (*model.Class, error) {
	if len(r.classes) == 0 {
		r.init()
	}

	for _, obj := range r.classes {
		if obj.ID == id {
			return obj, nil
		}
	}
	return nil, errors.New("no class wound with this id")
}

// GetAllClasses is the resolver for the getAllClasses field.
func (r *queryResolver) GetAllClasses(ctx context.Context) ([]*model.Class, error) {
	if len(r.classes) == 0 {
		r.init()
	}

	return r.classes, nil
}

// GetExerciseByID is the resolver for the getExerciseById field.
func (r *queryResolver) GetExerciseByID(ctx context.Context, id string) (*model.Exercise, error) {
	if len(r.exercises) == 0 {
		r.init()
	}
	for _, obj := range r.exercises {
		if obj.ID == id {
			return obj, nil
		}
	}
	return nil, errors.New("no exercise wound with this id")
}

// GetAllExercises is the resolver for the getAllExercises field.
func (r *queryResolver) GetAllExercises(ctx context.Context) ([]*model.Exercise, error) {
	if len(r.exercises) == 0 {
		r.init()
	}

	return r.exercises, nil
}

// GetAllResults is the resolver for the getAllResults field.
func (r *queryResolver) GetAllResults(ctx context.Context) ([]*model.Result, error) {
	if len(r.results) == 0 {
		r.init()
	}

	return r.results, nil
}

// GetLeaderBord is the resolver for the getLeaderBord field.
func (r *queryResolver) GetLeaderBord(ctx context.Context) ([]*model.LeaderboardRow, error) {
	if len(r.leaderboard) == 0 {
		r.init()
	}

	return r.leaderboard, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
