package gRPCClient

import (
	"Saga_orchestrator/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetAddressClient() (pb.GRPC_Address_ServerClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to dial: %v", err)
		return nil, nil, err
	}

	client := pb.NewGRPC_Address_ServerClient(conn)

	return client, conn, nil
}
