package auth

import (
	"errors"
)

type IPolicy interface {
	HasPermissions(bearerToken string, role string) bool
}

type Policy struct {
	Token ITokenProvider
}

func NewPolicy() IPolicy {
	token := NewToken()

	return &Policy{
		Token: token,
	}
}

func (p *Policy) HasPermissions(bearerToken string, role string) bool {
	_, roles, _ := p.getSubAndRoles(bearerToken)

	return p.hasRole(roles, role)
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
