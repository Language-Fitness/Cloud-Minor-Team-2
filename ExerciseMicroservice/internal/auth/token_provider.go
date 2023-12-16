package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ITokenProvider interface {
	IntrospectToken(bearerToken string) (bool, error)
	DecodeToken(token string) (map[string]interface{}, error)
	GenerateBasicAuthHeader() string
}

type TokenProvider struct {
	ClientID     string
	ClientSecret string
	Endpoint     string
}

func NewToken() ITokenProvider {
	clientID := getKeycloakClientId()
	clientSecret := getKeycloakClientSecret()
	endpoint := getKeycloakHost()

	tokenProvider := TokenProvider{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
	}

	return &tokenProvider
}

func (a *TokenProvider) IntrospectToken(bearerToken string) (bool, error) {
	authHeader := a.GenerateBasicAuthHeader()

	reqBody := fmt.Sprintf("token_type_hint=requesting_party_token&token=%s", bearerToken)
	req, err := http.NewRequest("POST", a.Endpoint, strings.NewReader(reqBody))
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var introspectionResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&introspectionResponse)
	if err != nil {
		return false, err
	}

	active, ok := introspectionResponse["active"].(bool)
	if !ok {
		return false, fmt.Errorf("unexpected response format: %v", introspectionResponse)
	}

	return active, nil
}

func (a *TokenProvider) DecodeToken(token string) (map[string]interface{}, error) {
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

func (a *TokenProvider) GenerateBasicAuthHeader() string {
	authString := fmt.Sprintf("%s:%s", a.ClientID, a.ClientSecret)
	authHeader := base64.StdEncoding.EncodeToString([]byte(authString))
	return fmt.Sprintf("Basic %s", authHeader)
}

func getKeycloakClientId() string {
	return os.Getenv("KEYCLOAK_CLIENT_ID")
}

func getKeycloakClientSecret() string {
	return os.Getenv("KEYCLOAK_CLIENT_SECRET")
}

func getKeycloakHost() string {
	return os.Getenv("KEYCLOAK_HOST") + "realms/" + os.Getenv("KEYCLOAK_REALM") + "/protocol/openid-connect/token/introspect"
}
