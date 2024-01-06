package service

import (
	"context"
	"school/proto/pb"
)

type SchoolGRPCService struct {
	pb.UnimplementedSchoolServiceServer
	service ISchoolService
}

func NewSchoolGRPCService() *SchoolGRPCService {
	return &SchoolGRPCService{
		service: NewSchoolService(),
	}
}

// GetKey implements the GetKey RPC method
func (s *SchoolGRPCService) GetKey(ctx context.Context, req *pb.KeyRequest) (*pb.KeyResponse, error) {
	response := &pb.KeyResponse{}

	// Get uuid from bearer token

	// Get user with that uuid

	// Get school and openai key with school_id of user

	school, err := s.service.GetSchoolById(req.BearerToken, req.SchoolId)
	if err != nil {
		response.Error = "school was not found"
		return response, nil
	}

	err = s.service.ValidateOpenAiKey(*school.OpenaiKey)
	if err != nil {
		response.Error = "key is not valid"
		return response, nil
	}

	response.Key = *school.OpenaiKey
	return response, nil
}
