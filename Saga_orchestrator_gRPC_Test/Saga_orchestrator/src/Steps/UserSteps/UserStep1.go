package UserSteps

import (
	"Saga_orchestrator/proto/pb"
	"Saga_orchestrator/src/gRPCClient"
	"context"
	"log"
)

const (
	UserStep1Execute    = "UserStep1Execute"
	UserStep1Compensate = "UserStep1Compensate"
)

func Step1Execute(Params []string) error {
	// Connect to gRPC service 1 and make a call
	client, conn, err := gRPCClient.GetUserClient()
	if err != nil {
		return err
	}
	defer conn.Close()

	req := &pb.GetUserRequest{UserId: Params[0]}
	user, err := client.GetUser(context.Background(), req)
	if err != nil {
		return err
	}
	log.Print(user.String())
	// Use the 'user' data if needed
	return nil
}

func Step1Compensate() error {
	// Implement the compensation logic for Step1
	log.Printf("Compensating Step, failed to get user")
	return nil
}
