package main

import (
	"Saga_orchestrator/src/Repository"
	"Saga_orchestrator/src/Saga"
	"Saga_orchestrator/src/Steps/AddressSteps"
	"Saga_orchestrator/src/Steps/UserSteps"
	"Saga_orchestrator/src/models"
	_ "github.com/google/uuid"
	"log"
)

func main() {
	// Init database
	Repository.InitMongoDB()

	step1 := models.SagaStep{
		Name:          "Step1 get user",
		Execute:       UserSteps.UserStep1Execute,
		Compensate:    UserSteps.UserStep1Compensate,
		ExecuteParams: []string{"1"},
	}

	step2 := models.SagaStep{
		Name:          "Step2 get address",
		Execute:       AddressSteps.AddressStep1Execute,
		Compensate:    AddressSteps.AddressStep1Compensate,
		ExecuteParams: []string{"1"},
	}

	// Create a new saga
	sagaID, err := Saga.CreateSaga(step1, step2)
	if err != nil {
		log.Fatalf("Failed to create saga: %v\n", err)
	}

	// Execute the saga
	if err := Saga.ExecuteSaga(sagaID); err != nil {
		log.Printf("Saga execution failed: %v\n", err)
	} else {
		log.Println("Saga completed successfully")
	}
}
