package grpcServers

import (
	"ResultMicroservice/proto/result"
	"context"
)

type ResultServer struct {
	result.UnimplementedGrpcResultServer
}

func (s *ResultServer) Delete(ctx context.Context, req *result.ResultRequest) (*result.Result, error) {
	//ToDO: Implement the Delete RPC
	result := &result.Result{
		//ToDo Import service method that uses
	}
	return result, nil
}

func (s *ResultServer) SoftDelete(ctx context.Context, req *result.ResultRequest) (*result.Result, error) {
	//ToDO: Implement the SoftDelete RPC
	result := &result.Result{
		//ToDo Import service method that uses
	}
	return result, nil
}
