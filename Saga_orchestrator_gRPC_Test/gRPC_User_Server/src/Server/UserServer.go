package Server

import (
	"context"
	"gRPC_User_Server/proto/pb"
	"gRPC_User_Server/src/Service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	pb.UnimplementedGRPC_User_ServerServer
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	// Replace with your logic to fetch user information from a data source
	user := Service.GetUserFromDataSource(req.UserId)

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}
