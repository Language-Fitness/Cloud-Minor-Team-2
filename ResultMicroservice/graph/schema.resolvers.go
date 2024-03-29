package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/auth"
	"context"
)

// CreateResult is the resolver for the CreateResult field.
func (r *mutationResolver) CreateResult(ctx context.Context, input model.InputResult) (*model.Result, error) {
	token := auth.TokenFromContext(ctx)

	result, err := r.Service.CreateResult(token, input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateResult is the resolver for the UpdateResult field.
func (r *mutationResolver) UpdateResult(ctx context.Context, id string, input model.InputResult) (*model.Result, error) {
	token := auth.TokenFromContext(ctx)

	result, err := r.Service.UpdateResult(token, id, input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ListResults is the resolver for the ListResults field.
func (r *queryResolver) ListResults(ctx context.Context, filter model.ResultFilter, paginator model.Paginator) ([]*model.ResultInfo, error) {
	token := auth.TokenFromContext(ctx)

	results, err := r.Service.ListResults(token, &filter, &paginator)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetResultsByID is the resolver for the GetResultsByID field.
func (r *queryResolver) GetResultsByID(ctx context.Context, id string) (*model.Result, error) {
	token := auth.TokenFromContext(ctx)

	result, err := r.Service.GetResultById(token, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
