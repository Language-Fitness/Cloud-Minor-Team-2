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
		service: NewModuleService(collection),
	}
}

// FindSagaObject implements the FindSagaObject RPC method
func (s *SagaService) FindSagaObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	module, err := s.service.GetModuleById(req.BearerToken, req.ObjectId)
	if err != nil {
		return nil, err
	}

	response := &pb.SagaObject{
		ObjectId:     module.ID,
		ObjectType:   pb.SagaObjectType_MODULE,
		ObjectStatus: setStatus(module.SoftDeleted),
	}

	return response, nil
}

// FindSagaObjectChildren implements the FindSagaObjectChildren RPC method
func (s *SagaService) FindSagaObjectChildren(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {

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
func (s *SagaService) DeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	err := s.service.DeleteModule(req.BearerToken, req.ObjectId, true)
	if err != nil {
		return nil, err
	}

	response := pb.SagaObject{
		ObjectId:     req.ObjectId,
		ObjectType:   pb.SagaObjectType_MODULE,
		ObjectStatus: pb.SagaObjectStatus_DELETED,
	}

	return &response, nil
}

// UnDeleteObject implements the UnDeleteObject RPC method
func (s *SagaService) UnDeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	err := s.service.DeleteModule(req.BearerToken, req.ObjectId, false)
	if err != nil {
		return nil, err
	}

	response := pb.SagaObject{
		ObjectId:     req.ObjectId,
		ObjectType:   pb.SagaObjectType_MODULE,
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
