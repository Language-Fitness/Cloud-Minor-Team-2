package service

import (
	"Module/proto/pb"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type SagaService struct {
	pb.UnimplementedGRPCSagaServiceServer
	service IModuleService
}

func NewSagaService(collection *mongo.Collection) *SagaService {
	return &SagaService{
		service: NewModuleService(collection), // Assuming you have a NewModuleService function
	}
}

// FindObject implements the FindObject RPC method
func (s *SagaService) FindObject(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {
	// Implement your logic to find the object based on the request
	// For demonstration purposes, let's just return a sample response
	response := &pb.ObjectResponse{
		Objects: []*pb.SagaObject{
			{
				ObjectId:     "sample_object_id",
				ObjectType:   pb.SagaObjectType_SCHOOL,
				ObjectStatus: pb.SagaObjectStatus_EXIST,
			},
		},
	}
	return response, nil
}

// DeleteObject implements the DeleteObject RPC method
func (s *SagaService) DeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {
	// Implement your logic to delete the object based on the request
	// For demonstration purposes, let's just return a sample response
	response := &pb.ObjectResponse{
		Objects: []*pb.SagaObject{
			{
				ObjectId:     "sample_object_id",
				ObjectType:   pb.SagaObjectType_SCHOOL,
				ObjectStatus: pb.SagaObjectStatus_DELETED,
			},
		},
	}
	return response, nil
}

// UnDeleteObject implements the UnDeleteObject RPC method
func (s *SagaService) UnDeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {
	// Implement your logic to undelete the object based on the request
	// For demonstration purposes, let's just return a sample response
	response := &pb.ObjectResponse{
		Objects: []*pb.SagaObject{
			{
				ObjectId:     "sample_object_id",
				ObjectType:   pb.SagaObjectType_SCHOOL,
				ObjectStatus: pb.SagaObjectStatus_EXIST,
			},
		},
	}
	return response, nil
}
