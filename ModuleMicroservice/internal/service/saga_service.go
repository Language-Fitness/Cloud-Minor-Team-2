package service

import (
	"Module/graph/model"
	"Module/internal/helper"
	"Module/proto/pb"
	"context"
	"fmt"
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

	filter := model.ModuleFilter{
		MadeBy: helper.StringPointer(req.ObjectId),
	}

	paginate := model.Paginator{
		Amount: 100,
		Step:   0,
	}

	modules, err := s.service.ListModules(req.BearerToken, helper.ModuleFilterPointer(filter), helper.PaginatorPointer(paginate))
	if err != nil {
		return nil, err
	}

	fmt.Println("test modules:")
	fmt.Println(modules)
	fmt.Println("test 1")
	fmt.Println("test 2")

	for i := range modules {
		fmt.Println(*modules[i])
	}

	response := &pb.ObjectResponse{
		Objects: make([]*pb.SagaObject, len(modules)),
	}

	for i := range modules {
		object := &pb.SagaObject{
			ObjectId:     fmt.Sprintf(modules[i].ID),
			ObjectType:   pb.SagaObjectType_MODULE,
			ObjectStatus: pb.SagaObjectStatus_EXIST,
		}

		response.Objects[i] = object
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
