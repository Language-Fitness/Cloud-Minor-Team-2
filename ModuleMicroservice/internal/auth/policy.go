package auth

import (
	"Module/graph/model"
	"Module/internal/repository"
	"context"
	"errors"
	"fmt"
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

func (p *Policy) CreateModule(bearerToken string, input model.ModuleInput) error {
	token, err := p.Token.IntrospectToken(bearerToken)
	if err != nil || token == false {
		return errors.New("invalid token")
	}

	decodeToken, err := p.Token.DecodeToken(bearerToken)
	if err != nil {
		return err
	}

	fmt.Println(decodeToken)

	sub, _ := decodeToken["sub"].(string)

	// Extract 'user-management-client' roles
	// Note: This assumes 'user-management-client' is always present in the map
	// and has a 'roles' key with a slice value.
	var roles []interface{}
	umcInterface, umcExists := decodeToken["user-management-client"]
	if umcExists {
		umc, ok := umcInterface.(map[string]interface{})
		if ok {
			rolesInterface, rolesExist := umc["roles"]
			if rolesExist {
				roles, _ = rolesInterface.([]interface{})
			}
		}
	}

	fmt.Println(sub, roles)

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
