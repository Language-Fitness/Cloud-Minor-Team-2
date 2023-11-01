package main

import (
	"gRPC_User_Server/proto/pb"
	"gRPC_User_Server/src/Server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051") // Define your gRPC server address and port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGRPC_User_ServerServer(server, &Server.UserServiceServer{})

	log.Println("Starting gRPC User Server on :50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
