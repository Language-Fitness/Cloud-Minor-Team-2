package service

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/helper"
	"ResultMicroservice/internal/service"
	"ResultMicroservice/proto/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type SagaService struct {
	pb.UnimplementedGRPCSagaServiceServer
	service service.IResultService
}

func NewSagaService(collection *mongo.Collection) *SagaService {
	return &SagaService{
		service: service.NewResultService(),
	}
}

// FindSagaObject implements the FindSagaObject RPC method
func (s *SagaService) FindSagaObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	result, err := s.service.GetResultById(req.BearerToken, req.ObjectId)
	if err != nil {
		return nil, err
	}

	response := &pb.SagaObject{
		ObjectId:     result.ID,
		ObjectType:   pb.SagaObjectType_RESULT,
		ObjectStatus: setStatus(result.SoftDeleted),
	}

	return response, nil
}

// FindSagaObjectChildren implements the FindSagaObjectChildren RPC method
func (s *SagaService) FindSagaObjectChildren(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {

	filter := model.ResultFilter{
		ExerciseID: helper.StringPointer(req.ObjectId),
	}

	paginate := model.Paginator{
		Amount: 100,
		Step:   0,
	}

	results, err := s.service.ListResults(req.BearerToken, &filter, &paginate)
	if err != nil {
		return nil, err
	}

	fmt.Println("test exercises:")
	fmt.Println(results)

	for i := range results {
		fmt.Println(*results[i])
	}

	response := &pb.ObjectResponse{
		Objects: make([]*pb.SagaObject, len(results)),
	}

	for i := range results {
		object := &pb.SagaObject{
			ObjectId:     fmt.Sprintf(results[i].ID),
			ObjectType:   pb.SagaObjectType_RESULT,
			ObjectStatus: pb.SagaObjectStatus_EXIST,
		}

		response.Objects[i] = object
	}

	return response, nil
}

// DeleteObject implements the DeleteObject RPC method
func (s *SagaService) DeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	err := s.service.DeleteResult(req.BearerToken, req.ObjectId, true)
	if err != nil {
		return nil, err
	}

	response := pb.SagaObject{
		ObjectId:     req.ObjectId,
		ObjectType:   pb.SagaObjectType_RESULT,
		ObjectStatus: pb.SagaObjectStatus_DELETED,
	}

	return &response, nil
}

// UnDeleteObject implements the UnDeleteObject RPC method
func (s *SagaService) UnDeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	err := s.service.DeleteResult(req.BearerToken, req.ObjectId, false)
	if err != nil {
		return nil, err
	}

	response := pb.SagaObject{
		ObjectId:     req.ObjectId,
		ObjectType:   pb.SagaObjectType_RESULT,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	return &response, nil
}

func setStatus(bool bool) pb.SagaObjectStatus {
	if bool == false {
		return pb.SagaObjectStatus_EXIST
	}

	return pb.SagaObjectStatus_DELETED
}
