package gRPCClient

import (
	"Saga_orchestrator/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetUserClient() (pb.GRPC_User_ServerClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewGRPC_User_ServerClient(conn)

	return client, conn, nil
}
