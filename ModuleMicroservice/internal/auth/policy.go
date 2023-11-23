package auth

import (
	"Module/graph/model"
	"Module/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Policy struct {
	Token            *Token
	ModuleRepository repository.IModuleRepository
}

func NewPolicy(collection *mongo.Collection) *Policy {
	token := NewToken()

	return &Policy{
		Token:            token,
		ModuleRepository: repository.NewModuleRepository(collection),
	}
}

func (p *Policy) CreateModule(ctx context.Context, input model.ModuleInput) (bool, error) {
	//headers := ctx.Value("headers").(http.Header)
	//
	//// Access tokens from the headers, e.g., for Bearer token
	//accessToken := headers.Get("Authorization")

	return true, nil
}

// UpdateModule is the resolver for the updateModule field.
func (p *Policy) UpdateModule(ctx context.Context, id string, input model.ModuleInput) (bool, error) {

	return true, nil
}

// DeleteModule is the resolver for the deleteModule field.
func (p *Policy) DeleteModule(ctx context.Context, id string) (bool, error) {

	return true, nil
}

// GetModule is the resolver for the getModule field.
func (p *Policy) GetModule(ctx context.Context, id string) (bool, error) {

	return true, nil
}

// ListModules is the resolver for the listModules field.
func (p *Policy) ListModules(ctx context.Context) (bool, error) {

	return true, nil
}
