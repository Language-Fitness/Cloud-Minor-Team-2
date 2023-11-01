package main

import (
	"log"
	"net"

	"gRPC_Address_Server/proto/pb" // Import the generated protobuf package
	"gRPC_Address_Server/src/Server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052") // Define your gRPC server address and port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGRPC_Address_ServerServer(server, &Server.AddressServiceServer{})

	log.Println("Starting gRPC Address Server on :50052...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
