package grpcServers

import (
	"ResultMicroservice/proto/result_pb"
	"context"
)

type ResultServer struct {
	result_pb.UnimplementedGrpcResultServer
}

func (s *ResultServer) DeleteByModule(ctx context.Context, req *result_pb.DeleteByModuleRequest) (*result_pb.Result, error) {
	// TODO: Implement the DeleteByModule RPC
	result := &result_pb.Result{
		// TODO: Call the service method that uses req.ModuleID and req.BearerToken
	}
	return result, nil
}

func (s *ResultServer) SoftDeleteByModule(ctx context.Context, req *result_pb.DeleteByModuleRequest) (*result_pb.Result, error) {
	// TODO: Implement the SoftDeleteByModule RPC
	result := &result_pb.Result{
		// TODO: Call the service method that uses req.ModuleID and req.BearerToken
	}
	return result, nil
}

func (s *ResultServer) DeleteByClass(ctx context.Context, req *result_pb.DeleteByClassRequest) (*result_pb.Result, error) {
	// TODO: Implement the DeleteByClass RPC
	result := &result_pb.Result{
		// TODO: Call the service method that uses req.ClassID and req.BearerToken
	}
	return result, nil
}

func (s *ResultServer) SoftDeleteByClass(ctx context.Context, req *result_pb.DeleteByClassRequest) (*result_pb.Result, error) {
	// TODO: Implement the SoftDeleteByClass RPC
	result := &result_pb.Result{
		// TODO: Call the service method that uses req.ClassID and req.BearerToken
	}
	return result, nil
}

func (s *ResultServer) DeleteByUser(ctx context.Context, req *result_pb.DeleteByUserRequest) (*result_pb.Result, error) {
	// TODO: Implement the DeleteByUser RPC
	result := &result_pb.Result{
		// TODO: Call the service method that uses req.UserID and req.BearerToken
	}
	return result, nil
}

func (s *ResultServer) SoftDeleteByUser(ctx context.Context, req *result_pb.DeleteByUserRequest) (*result_pb.Result, error) {
	// TODO: Implement the SoftDeleteByUser RPC
	result := &result_pb.Result{
		// TODO: Call the service method that uses req.UserID and req.BearerToken
	}
	return result, nil
}
