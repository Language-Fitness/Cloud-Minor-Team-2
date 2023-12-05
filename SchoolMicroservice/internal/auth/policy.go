package auth

import (
	"errors"
	"example/graph/model"
	"example/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPolicy interface {
	CreateSchool(bearerToken string) (string, error)
	UpdateSchool(bearerToken string, id string) (*model.School, error)
	DeleteSchool(bearerToken string, id string) (*model.School, error)
	GetSchool(bearerToken string, id string) (*model.School, error)
	ListSchools(bearerToken string) error
}

type Policy struct {
	Token            IToken
	SchoolRepository repository.ISchoolRepository
}

func NewPolicy(collection *mongo.Collection) IPolicy {
	token := NewToken()

	return &Policy{
		Token:            token,
		SchoolRepository: repository.NewSchoolRepository(collection),
	}
}

func (p *Policy) CreateSchool(bearerToken string) (string, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return "", err2
	}

	if !p.hasRole(roles, "create_school") {
		return "", errors.New("invalid permissions for this action")
	}

	return uuid, nil
}

func (p *Policy) UpdateSchool(bearerToken string, id string) (*model.School, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	school, err := p.SchoolRepository.GetSchoolByID(id)
	if err != nil {
		return nil, errors.New("invalid permissions for this action")
	}

	if p.hasRole(roles, "update_school") && school.MadeBy == uuid {
		return school, nil
	}

	if p.hasRole(roles, "update_school_all") {
		return school, nil
	}

	return nil, errors.New("invalid permissions for this action")
}

func (p *Policy) DeleteSchool(bearerToken string, id string) (*model.School, error) {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	school, err := p.SchoolRepository.GetSchoolByID(id)
	if err != nil {
		return nil, errors.New("invalid permissions for this action")
	}

	if p.hasRole(roles, "delete_school_all") {
		return school, nil
	}

	return nil, errors.New("invalid permissions for this action")
}

func (p *Policy) GetSchool(bearerToken string, id string) (*model.School, error) {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	school, err := p.SchoolRepository.GetSchoolByID(id)
	if err != nil {
		return nil, errors.New("invalid permissions for this action")
	}

	if p.hasRole(roles, "get_school") {
		return school, nil
	}

	return nil, errors.New("invalid permissions for this action")
}

func (p *Policy) ListSchools(bearerToken string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !p.hasRole(roles, "get_schools") {
		return errors.New("invalid permissions for this action")
	}

	return nil
}

func (p *Policy) getSubAndRoles(bearerToken string) (string, []interface{}, error) {
	//token, err := p.Token.IntrospectToken(bearerToken)
	//if err != nil || token == false {
	//	return "", nil, errors.New("invalid token")
	//}

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
