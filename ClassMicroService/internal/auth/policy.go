package auth

import (
	"Class/graph/model"
	"Class/internal/repository"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

const InvalidActionsMessage = "invalid permissions for this action"
const InvalidTokenMessage = "invalid token"

type IPolicy interface {
	CreateClass(bearerToken string) (string, error)
	UpdateClass(bearerToken string, id string) (*model.Class, error)
	DeleteClass(bearerToken string, id string) (*model.Class, error)
	GetClass(bearerToken string, id string) (*model.Class, error)
	ListClasses(bearerToken string) error
	HasPermissions(bearerToken string, role string) bool
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
		return "", errors.New(InvalidActionsMessage)
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
		return nil, errors.New(InvalidActionsMessage)
	}

	if p.hasRole(roles, "update_class") && class.MadeBy == uuid {
		return class, nil
	}

	if p.hasRole(roles, "update_class_all") {
		return class, nil
	}

	return nil, errors.New(InvalidActionsMessage)
}

func (p *Policy) DeleteClass(bearerToken string, id string) (*model.Class, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	class, err := p.ClassRepository.GetClassByID(id)
	if err != nil {
		return nil, errors.New("class not found")
	}

	if p.hasRole(roles, "delete_class_all") {
		return class, nil
	}

	if p.hasRole(roles, "delete_class") && class.MadeBy == uuid {
		return class, nil
	}

	return nil, errors.New(InvalidActionsMessage)
}

func (p *Policy) GetClass(bearerToken string, id string) (*model.Class, error) {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	class, err := p.ClassRepository.GetClassByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(InvalidActionsMessage)
	}

	if p.hasRole(roles, "get_class") {
		fmt.Println("does have role")
		return class, nil
	}

	fmt.Println("does not have role")
	return nil, errors.New(InvalidActionsMessage)
}

func (p *Policy) ListClasses(bearerToken string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if p.hasRole(roles, "get_classes_all") {
		return nil
	}

	if !p.hasRole(roles, "get_classes") {
		return errors.New(InvalidActionsMessage)
	}

	return nil
}

func (p *Policy) HasPermissions(bearerToken string, role string) bool {
	_, roles, _ := p.getSubAndRoles(bearerToken)

	return p.hasRole(roles, role)
}

func (p *Policy) getSubAndRoles(bearerToken string) (string, []interface{}, error) {
	//token, err := p.Token.IntrospectToken(bearerToken)
	//if err != nil || token == false {
	//	return "", nil, errors.New("invalid token introspect")
	//}

	decodeToken, err := p.Token.DecodeToken(bearerToken)
	if err != nil {
		return "", nil, err
	}

	sub, _ := decodeToken["sub"].(string)

	resourceAccess, ok := decodeToken["resource_access"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New(InvalidTokenMessage)
	}

	userManagementClient, ok := resourceAccess["user-management-client"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New(InvalidTokenMessage)
	}

	roles, ok := userManagementClient["roles"].([]interface{})
	if !ok {
		return "", nil, errors.New(InvalidTokenMessage)
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
