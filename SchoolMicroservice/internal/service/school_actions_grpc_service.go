package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"school/internal/auth"
	"school/proto/pb"
)

type SchoolActionsGRPCService struct {
	pb.UnimplementedSchoolServiceServer
	service ISchoolService
}

func NewSchoolGRPCService() *SchoolActionsGRPCService {
	return &SchoolActionsGRPCService{
		service: NewSchoolService(),
	}
}

// GetKey implements the GetKey RPC method
func (s *SchoolActionsGRPCService) GetKey(ctx context.Context, req *pb.KeyRequest) (*pb.KeyResponse, error) {
	response := &pb.KeyResponse{}

	// get user uuid from bearer
	uuid, err := getSub(req.BearerToken)
	if err != nil {
		response.Error = err.Error()
		return response, nil
	}

	// retrieve schoolId from keycloak with uuid
	schoolId, err := getSchoolId(req.BearerToken, uuid)
	if err != nil {
		response.Error = err.Error()
		return response, nil
	}

	// retrieve school from db using school id in user
	school, err := s.service.GetSchoolById(req.BearerToken, schoolId)
	if err != nil {
		response.Error = "school was not found"
		return response, nil
	}

	// validate retrieved key using test request
	err = s.service.ValidateOpenAiKey(*school.OpenaiKey)
	if err != nil {
		response.Error = "key is not valid"
		return response, nil
	}

	response.Key = *school.OpenaiKey
	return response, nil
}

func getSub(bearerToken string) (string, error) {
	token := auth.NewToken()
	decodeToken, err := token.DecodeToken(bearerToken)
	if err != nil {
		return "", err
	}

	sub, _ := decodeToken["sub"].(string)
	return sub, nil
}

func getSchoolId(token string, userId string) (string, error) {
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

	schoolId, err := extractSchoolIDFromJson(result)
	if err != nil {
		return "", err
	}

	return schoolId, nil
}

func extractSchoolIDFromJson(result map[string]interface{}) (string, error) {
	attributes, ok := result["attributes"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to extract attributes")
	}

	schoolIDList, ok := attributes["school_id"].([]interface{})
	if !ok || len(schoolIDList) == 0 {
		return "", fmt.Errorf("school_id not found or empty")
	}

	schoolID, ok := schoolIDList[0].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract school_id")
	}

	return schoolID, nil
}
