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
func (s *SagaService) FindSagaObject(context context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	class, err := s.service.GetClassById(req.BearerToken, req.ObjectId)
	if err != nil {
		return nil, err
	}

	response := &pb.SagaObject{
		ObjectId:     class.ID,
		ObjectType:   pb.SagaObjectType_CLASS,
		ObjectStatus: setStatus(class.SoftDeleted),
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
func (s *SagaService) DeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	err := s.service.DeleteClass(req.BearerToken, req.ObjectId, true)
	if err != nil {
		return nil, err
	}

	fmt.Println("i execute this")

	response := pb.SagaObject{
		ObjectId:     req.ObjectId,
		ObjectType:   pb.SagaObjectType_CLASS,
		ObjectStatus: pb.SagaObjectStatus_DELETED,
	}

	return &response, nil
}

// UnDeleteObject implements the UnDeleteObject RPC method
func (s *SagaService) UnDeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	err := s.service.DeleteClass(req.BearerToken, req.ObjectId, false)
	if err != nil {
		return nil, err
	}

	response := pb.SagaObject{
		ObjectId:     req.ObjectId,
		ObjectType:   pb.SagaObjectType_CLASS,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	return &response, nil
}

func setStatus(bool *bool) pb.SagaObjectStatus {
	if bool == helper.BoolPointer(false) {
		return pb.SagaObjectStatus_EXIST
	}

	return pb.SagaObjectStatus_DELETED
}
