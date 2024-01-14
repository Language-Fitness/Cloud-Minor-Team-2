package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"school/internal/helper"
	"school/proto/pb"
)

type SagaService struct {
	pb.UnimplementedGRPCSagaServiceServer
	service ISchoolService
}

func NewSagaService(collection *mongo.Collection) *SagaService {
	return &SagaService{
		service: NewSchoolService(),
	}
}

// FindSagaObject implements the FindSagaObject RPC method
func (s *SagaService) FindSagaObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	sagaObject, err := s.service.GetSchoolById(req.BearerToken, req.ObjectId)
	if err != nil {
		return nil, err
	}

	response := &pb.SagaObject{
		ObjectId:     sagaObject.ID,
		ObjectType:   pb.SagaObjectType_SCHOOL,
		ObjectStatus: setStatus(sagaObject.SoftDeleted),
	}

	return response, nil
}

// FindSagaObjectChildren implements the FindSagaObjectChildren RPC method
func (s *SagaService) FindSagaObjectChildren(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {

	//No need for implementations because this can never be a child

	return nil, nil
}

// DeleteObject implements the DeleteObject RPC method
func (s *SagaService) DeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	err := s.service.DeleteSchool(req.BearerToken, req.ObjectId, true)
	if err != nil {
		return nil, err
	}

	response := pb.SagaObject{
		ObjectId:     req.ObjectId,
		ObjectType:   pb.SagaObjectType_SCHOOL,
		ObjectStatus: pb.SagaObjectStatus_DELETED,
	}

	return &response, nil
}

// UnDeleteObject implements the UnDeleteObject RPC method
func (s *SagaService) UnDeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	err := s.service.DeleteSchool(req.BearerToken, req.ObjectId, false)
	if err != nil {
		return nil, err
	}

	response := pb.SagaObject{
		ObjectId:     req.ObjectId,
		ObjectType:   pb.SagaObjectType_SCHOOL,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	return &response, nil
}

func setStatus(bool *bool) pb.SagaObjectStatus {
	if bool == helper.BoolPointer(true) {
		return pb.SagaObjectStatus_EXIST
	}

	return pb.SagaObjectStatus_DELETED
}
