package main

import (
	"context"
	"gRPC-Service/src/services"
	"log"
	"net"

	"gRPC-Service/proto/pb"
	"google.golang.org/grpc"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	log.Println("GetAllUsers")
	// Create a slice of pointers to pb.User.
	var userPointers, _ = services.GetAllUsers()

	// Create the response with the slice of user pointers.
	response := &pb.GetAllUsersResponse{
		Users: userPointers,
	}

	return response, nil
}

func (s *UserServiceServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	log.Println("GetUserById")
	return services.GetUserByID(int(req.UserId))
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServiceServer{})

	log.Println("Starting gRPC server on :50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
