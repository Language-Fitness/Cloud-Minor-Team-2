package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// IUserInfoProvider interface for dependency injection
type IUserInfoProvider interface {
	GetUserInfo(token string, userId string) (string, error)
}

// DefaultUserInfoProvider implements UserInfoProvider
type DefaultUserInfoProvider struct{}

// NewUserProvider GOLANG FACTORY
// Returns a DefaultUserInfoProvider implementing IUserInfoProvider.
func NewUserProvider() IUserInfoProvider {
	return &DefaultUserInfoProvider{}
}

func (d *DefaultUserInfoProvider) GetUserInfo(token string, userId string) (string, error) {
	baseUrl := os.Getenv("KEYCLOAK_HOST")
	path := "admin/realms/cloud-project/users/"
	fullUrl := baseUrl + path + userId

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	fullName, err := extractNamesFromJson(result)
	if err != nil {
		return "", err
	}

	return fullName, nil
}

func extractNamesFromJson(result map[string]interface{}) (string, error) {
	firstName, ok := result["firstName"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract firstName")
	}

	lastName, ok := result["lastName"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract lastName")
	}

	fullName := firstName + " " + lastName
	return fullName, nil
}
