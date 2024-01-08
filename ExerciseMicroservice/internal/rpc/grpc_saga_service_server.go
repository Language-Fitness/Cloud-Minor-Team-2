package service

import (
	"ExerciseMicroservice/graph/model"
	"ExerciseMicroservice/internal/helper"
	"ExerciseMicroservice/internal/service"
	"ExerciseMicroservice/proto/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type SagaService struct {
	pb.UnimplementedGRPCSagaServiceServer
	service service.IExerciseService
}

func NewSagaService(collection *mongo.Collection) *SagaService {
	return &SagaService{
		service: service.NewExerciseService(),
	}
}

// FindSagaObject implements the FindSagaObject RPC method
func (s *SagaService) FindSagaObject(ctx context.Context, req *pb.ObjectRequest) (*pb.SagaObject, error) {
	result, err := s.service.GetExerciseById(req.BearerToken, req.ObjectId)
	if err != nil {
		return nil, err
	}

	response := &pb.SagaObject{
		ObjectId:     result.ID,
		ObjectType:   pb.SagaObjectType_EXERCISE,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	return response, nil
}

// FindSagaObjectChildren implements the FindSagaObjectChildren RPC method
func (s *SagaService) FindSagaObjectChildren(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {

	filter := model.ExerciseFilter{
		ClassID: helper.StringPointer(req.ObjectId),
	}

	paginate := model.Paginator{
		Amount: 100,
		Step:   0,
	}

	exercises, err := s.service.ListExercises(req.BearerToken, &filter, &paginate)
	if err != nil {
		return nil, err
	}

	fmt.Println("test exercises:")
	fmt.Println(exercises)

	for i := range exercises {
		fmt.Println(*exercises[i])
	}

	response := &pb.ObjectResponse{
		Objects: make([]*pb.SagaObject, len(exercises)),
	}

	for i := range exercises {
		object := &pb.SagaObject{
			ObjectId:     fmt.Sprintf(exercises[i].ID),
			ObjectType:   pb.SagaObjectType_EXERCISE,
			ObjectStatus: pb.SagaObjectStatus_EXIST,
		}

		response.Objects[i] = object
	}

	return response, nil
}

// DeleteObject implements the DeleteObject RPC method
func (s *SagaService) DeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {
	err := s.service.DeleteExercise(req.BearerToken, req.ObjectId)
	if err != nil {
		return nil, err
	}

	response := &pb.ObjectResponse{
		Objects: []*pb.SagaObject{
			{
				ObjectId:     "req.ObjectId",
				ObjectType:   pb.SagaObjectType_EXERCISE,
				ObjectStatus: pb.SagaObjectStatus_DELETED,
			},
		},
	}
	return response, nil
}

// UnDeleteObject implements the UnDeleteObject RPC method
func (s *SagaService) UnDeleteObject(ctx context.Context, req *pb.ObjectRequest) (*pb.ObjectResponse, error) {
	err := s.service.UnDeleteExercise(req.BearerToken, req.ObjectId)
	if err != nil {
		return nil, err
	}

	response := &pb.ObjectResponse{
		Objects: []*pb.SagaObject{
			{
				ObjectId:     req.ObjectId,
				ObjectType:   pb.SagaObjectType_EXERCISE,
				ObjectStatus: pb.SagaObjectStatus_EXIST,
			},
		},
	}
	return response, nil
}
