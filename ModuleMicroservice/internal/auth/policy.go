package auth

import (
	"Module/internal/repository"
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

func (p *Policy) UpdateModule(bearerToken string, id string) error {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	module, err := p.ModuleRepository.GetModuleByID(id)
	if err != nil {
		return errors.New("invalid permissions for this action")
	}

	if hasRole(roles, "update_module") && *module.MadeBy == uuid {
		return nil
	}

	if hasRole(roles, "update_module_all") {
		return nil
	}

	return errors.New("invalid permissions for this action")
}

func (p *Policy) DeleteModule(bearerToken string, id string) error {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	module, err := p.ModuleRepository.GetModuleByID(id)
	if err != nil {
		return errors.New("invalid permissions for this action")
	}

	if hasRole(roles, "delete_module") && *module.MadeBy == uuid {
		return nil
	}

	if hasRole(roles, "delete_module_all") {
		return nil
	}

	return errors.New("invalid permissions for this action")
}

// GetModule is the resolver for the getModule field.
func (p *Policy) GetModule(bearerToken string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !hasRole(roles, "get_module") {
		return errors.New("invalid permissions for this action")
	}

	return nil
}

// ListModules is the resolver for the listModules field.
func (p *Policy) ListModules(bearerToken string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !hasRole(roles, "get_modules") {
		return errors.New("invalid permissions for this action")
	}

	return nil
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

	userManagementClient, ok := resourceAccess["user-management-client"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New("invalid token")
	}

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
