package service

import (
	"Class/graph/model"
	"Class/internal/helper"
	"Class/proto/pb"
	"context"
	"fmt"
)

type SagaService struct {
	pb.UnimplementedGRPCSagaServiceServer
	service IClassService
}

func NewSagaService() *SagaService {
	return &SagaService{
		service: NewClassService(),
	}
}

// FindSagaObject implements the FindSagaObject RPC method
func (s *SagaService) FindSagaObject(context.Context, *pb.ObjectRequest) (*pb.SagaObject, error) {
	response := &pb.SagaObject{
		ObjectId:     "0e520bea-a96b-47cc-96bc-83633e47c58e",
		ObjectType:   pb.SagaObjectType_SCHOOL,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	return response, nil
}

// FindSagaObjectChildren implements the FindSagaObjectChildren RPC method
func (s *SagaService) FindSagaObjectChildren(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {

	filter := model.ListClassFilter{
		ModuleID: helper.StringPointer(req.ObjectId),
	}

	paginate := model.Paginator{
		Amount: 100,
		Step:   0,
	}

	fmt.Println(filter)

	classes, err := s.service.ListClasses(req.BearerToken, helper.ListClassFilterPointer(filter), helper.PaginatorPointer(paginate))
	if err != nil {
		return nil, err
	}

	fmt.Println(classes)

	for i := range classes {
		fmt.Println(*classes[i])
	}

	response := &pb.ObjectResponse{
		Objects: make([]*pb.SagaObject, len(classes)),
	}

	for i := range classes {
		object := &pb.SagaObject{
			ObjectId:     fmt.Sprintf(classes[i].ID),
			ObjectType:   pb.SagaObjectType_CLASS,
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
