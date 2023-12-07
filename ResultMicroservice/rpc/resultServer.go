package rpc

import (
	"ResultMicroservice/internal/service"
	"ResultMicroservice/proto/result_pb"
	"context"
)

type ResultServer struct {
	result_pb.UnimplementedGrpcResultServer
	ResultService service.IResultService
}

func NewResultServer(resultService service.IResultService) *ResultServer {
	return &ResultServer{
		ResultService: resultService,
	}
}

func (s *ResultServer) DeleteByModule(ctx context.Context, req *result_pb.DeleteByModuleRequest) (*result_pb.Response, error) {
	id, deleted, err := s.ResultService.DeleteByModule(req.BearerToken, req.ModuleID)
	if err != nil {

		return nil, err
	}

	return &result_pb.Response{ID: id, Deleted: deleted}, nil
}

func (s *ResultServer) SoftDeleteByModule(ctx context.Context, req *result_pb.DeleteByModuleRequest) (*result_pb.Response, error) {
	id, deleted, err := s.ResultService.SoftDeleteByModule(req.BearerToken, req.ModuleID)
	if err != nil {

		return nil, err
	}

	return &result_pb.Response{ID: id, Deleted: deleted}, nil
}

func (s *ResultServer) DeleteByClass(ctx context.Context, req *result_pb.DeleteByClassRequest) (*result_pb.Response, error) {
	id, deleted, err := s.ResultService.DeleteByClass(req.BearerToken, req.ClassID)
	if err != nil {

		return nil, err
	}

	return &result_pb.Response{ID: id, Deleted: deleted}, nil
}

func (s *ResultServer) SoftDeleteByClass(ctx context.Context, req *result_pb.DeleteByClassRequest) (*result_pb.Response, error) {
	id, deleted, err := s.ResultService.SoftDeleteByClass(req.BearerToken, req.ClassID)
	if err != nil {

		return nil, err
	}

	return &result_pb.Response{ID: id, Deleted: deleted}, nil
}

func (s *ResultServer) DeleteByUser(ctx context.Context, req *result_pb.DeleteByUserRequest) (*result_pb.Response, error) {
	id, deleted, err := s.ResultService.DeleteByUser(req.BearerToken, req.UserID)
	if err != nil {

		return nil, err
	}

	return &result_pb.Response{ID: id, Deleted: deleted}, nil
}

func (s *ResultServer) SoftDeleteByUser(ctx context.Context, req *result_pb.DeleteByUserRequest) (*result_pb.Response, error) {
	id, deleted, err := s.ResultService.SoftDeleteByUser(req.BearerToken, req.UserID)
	if err != nil {

		return nil, err
	}

	return &result_pb.Response{ID: id, Deleted: deleted}, nil
}
