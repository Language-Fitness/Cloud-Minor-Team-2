package mocks

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/mock"
	"strings"
)

type MockToken struct {
	mock.Mock
	ClientID     string
	ClientSecret string
	Endpoint     string
}

func (m *MockToken) IntrospectToken(bearerToken string) (bool, error) {
	args := m.Called(bearerToken)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockToken) DecodeToken(token string) (map[string]interface{}, error) {
	// JWTs are typically in the format "header.payload.signature"
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	// Decode the payload (second part)
	decodedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("error decoding token payload: %v", err)
	}

	// Unmarshal the JSON payload into a map
	var claims map[string]interface{}
	err = json.Unmarshal(decodedPayload, &claims)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling token claims: %v", err)
	}

	return claims, nil
}

func (m *MockToken) GenerateBasicAuthHeader() string {
	args := m.Called()
	return args.Get(0).(string)
}
