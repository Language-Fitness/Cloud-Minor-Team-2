package auth

import (
	"Module/graph/model"
	"Module/internal/repository"
	"context"
	"errors"
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

func (p *Policy) CreateModule(bearerToken string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !hasRole(roles, "create_module") {
		return errors.New("invalid permissions for this action")
	}

	return nil
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

func (p *Policy) getSubAndRoles(bearerToken string) (string, []interface{}, error) {
	token, err := p.Token.IntrospectToken(bearerToken)
	if err != nil || token == false {
		return "", nil, errors.New("invalid token")
	}

	decodeToken, err := p.Token.DecodeToken(bearerToken)
	if err != nil {
		return "", nil, err
	}

	sub, _ := decodeToken["sub"].(string)

	resourceAccess, ok := decodeToken["resource_access"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New("invalid token")
	}

	// Access the 'user-management-client' map
	userManagementClient, ok := resourceAccess["user-management-client"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New("invalid token")
	}

	// Access the 'roles' array within 'user-management-client'
	roles, ok := userManagementClient["roles"].([]interface{})
	if !ok {
		return "", nil, errors.New("invalid token")
	}
	return sub, roles, nil
}

func hasRole(roles []interface{}, targetRole string) bool {
	for _, role := range roles {
		if role == targetRole {
			return true
		}
	}
	return false
}
