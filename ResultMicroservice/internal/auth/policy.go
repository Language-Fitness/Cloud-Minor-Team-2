package auth

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/repository"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const InvalidTokenMessage = "invalid token"

type IResultPolicy interface {
	CreateResult(bearerToken string) (string, error)
	UpdateResult(bearerToken string, id string) (*model.Result, error)
	DeleteResult(bearerToken string, id string) (*model.Result, error)
	GetResultByID(bearerToken string, id string) (*model.Result, error)
	ListResult(bearerToken string) (string, bool, error)
	HasPermissions(bearerToken string, role string) bool
}

type ResultPolicy struct {
	Token            IToken
	ResultRepository repository.IResultRepository
}

func NewResultPolicy(collection *mongo.Collection) IResultPolicy {
	token := NewToken()

	return &ResultPolicy{
		Token:            token,
		ResultRepository: repository.NewResultRepository(collection),
	}
}

func (p *ResultPolicy) CreateResult(bearerToken string) (string, error) {
	uuid, roles, err := p.getSubAndRoles(bearerToken)
	if err != nil {
		return "", err
	}

	if !p.hasRole(roles, "create_result") {
		return "", errors.New("invalid permissions for this action")
	}

	return uuid, nil
}

func (p *ResultPolicy) UpdateResult(bearerToken string, id string) (*model.Result, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	result, err := p.ResultRepository.GetResultByID(id)
	if err != nil {
		return nil, errors.New("result not found")
	}

	if p.hasRole(roles, "update_result_all") {
		return result, nil
	}

	if p.hasRole(roles, "update_result") && result.UserID == uuid {
		return result, nil
	}

	return nil, errors.New("invalid permissions for this action")
}

func (p *ResultPolicy) DeleteResult(bearerToken string, id string) (*model.Result, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	result, err := p.ResultRepository.GetResultByID(id)
	if err != nil {
		return nil, errors.New("result not found")
	}

	if p.hasRole(roles, "delete_result") && result.UserID == uuid {
		return result, nil
	}

	if p.hasRole(roles, "delete_result_all") {
		return result, nil
	}

	return nil, errors.New("invalid permissions for this action")
}

func (p *ResultPolicy) GetResultByID(bearerToken string, id string) (*model.Result, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	result, err := p.ResultRepository.GetResultByID(id)
	if err != nil {
		return nil, errors.New("result not found")
	}

	if p.hasRole(roles, "get_result") && result.UserID == uuid {
		return result, nil
	}

	if p.hasRole(roles, "get_result_all") {
		return result, nil
	}

	return nil, errors.New("invalid permissions for this action")
}

func (p *ResultPolicy) ListResult(bearerToken string) (string, bool, error) {
	uuid, roles, err := p.getSubAndRoles(bearerToken)
	if err != nil {
		return "", false, err
	}

	if p.hasRole(roles, "list_results_all") {
		return uuid, true, nil
	}

	if p.hasRole(roles, "list_results") {
		return uuid, false, nil
	}

	return uuid, false, errors.New("invalid permissions for this action")
}

func (p *ResultPolicy) HasPermissions(bearerToken string, role string) bool {
	_, roles, _ := p.getSubAndRoles(bearerToken)

	return p.hasRole(roles, role)
}

func (p *ResultPolicy) getSubAndRoles(bearerToken string) (string, []interface{}, error) {
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

func (p *ResultPolicy) hasRole(roles []interface{}, targetRole string) bool {
	for _, role := range roles {
		if role == targetRole {
			return true
		}
	}
	return false
}
