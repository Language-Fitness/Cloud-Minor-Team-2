package AddressSteps

import (
	"Saga_orchestrator/proto/pb"
	"Saga_orchestrator/src/gRPCClient"
	"context"
	"log"
)

type AddressStepFunctions string

const (
	AddressStep1Execute    = "AddressStep1Execute"
	AddressStep1Compensate = "AddressStep1Compensate"
)

func Step1Execute(Params []string) error {
	// Connect to gRPC Address service and make a call
	client, conn, err := gRPCClient.GetAddressClient()
	if err != nil {
		return err
	}
	defer conn.Close()

	req := &pb.GetUserAddressRequest{UserId: Params[0]}
	address, err := client.GetUserAddress(context.Background(), req)
	if err != nil {
		return err
	}
	log.Print(address.Address)

	// Use the 'address' data if needed
	return nil
}

func Step1Compensate() error {
	// Implement the compensation logic for Step1
	log.Printf("Compensating Step, failed to get address")
	return nil
}
