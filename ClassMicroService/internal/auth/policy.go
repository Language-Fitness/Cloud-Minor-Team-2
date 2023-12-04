package auth

import (
	"errors"
	"example/graph/model"
	"example/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPolicy interface {
	CreateClass(bearerToken string) (string, error)
	UpdateClass(bearerToken string, id string) (*model.Class, error)
	DeleteClass(bearerToken string, id string) error
	GetClass(bearerToken string) error
	ListClasses(bearerToken string) error
}

type Policy struct {
	Token           IToken
	ClassRepository repository.IClassRepository
}

func NewPolicy(collection *mongo.Collection) IPolicy {
	token := NewToken()

	return &Policy{
		Token:           token,
		ClassRepository: repository.NewClassRepository(collection),
	}
}

func (p *Policy) CreateClass(bearerToken string) (string, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return "", err2
	}

	if !p.hasRole(roles, "create_class") {
		return "", errors.New("invalid permissions for this action")
	}

	return uuid, nil
}

func (p *Policy) UpdateClass(bearerToken string, id string) (*model.Class, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	class, err := p.ClassRepository.GetClassByID(id)
	if err != nil {
		return nil, errors.New("invalid permissions for this action")
	}

	if p.hasRole(roles, "update_class") && class.MadeBy == uuid {
		return class, nil
	}

	if p.hasRole(roles, "update_class_all") {
		return class, nil
	}

	return nil, errors.New("invalid permissions for this action")
}

func (p *Policy) DeleteClass(bearerToken string, id string) error {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	class, err := p.ClassRepository.GetClassByID(id)
	if err != nil {
		return errors.New("invalid permissions for this action")
	}

	if p.hasRole(roles, "delete_class") && class.MadeBy == uuid {
		return nil
	}

	if p.hasRole(roles, "delete_class_all") {
		return nil
	}

	return errors.New("invalid permissions for this action")
}

func (p *Policy) GetClass(bearerToken string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !p.hasRole(roles, "get_class") {
		return errors.New("invalid permissions for this action")
	}

	return nil
}

func (p *Policy) ListClasses(bearerToken string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !p.hasRole(roles, "get_classes") {
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

func (p *Policy) hasRole(roles []interface{}, targetRole string) bool {
	for _, role := range roles {
		if role == targetRole {
			return true
		}
	}
	return false
}
