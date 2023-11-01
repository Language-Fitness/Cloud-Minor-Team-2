package main

import (
	AddressSteps "Saga_orchestrator/src/AddressSteps"
	"Saga_orchestrator/src/Saga"
	"Saga_orchestrator/src/UserSteps"
	"log"
)

func main() {

	step1 := Saga.SagaStep{
		Name: "Step1",
		Execute: func() error {
			return UserSteps.Step1Execute("1")
		},
		Compensate: UserSteps.Step1Compensate,
	}

	step2 := Saga.SagaStep{
		Name: "Step2",
		Execute: func() error {
			return AddressSteps.Step1Execute("1")
		},
		Compensate: AddressSteps.Step1Compensate,
	}
	// Create a new saga
	sagaID := Saga.CreateSaga(step1, step2)

	// Execute the saga
	if err := Saga.ExecuteSaga(sagaID); err != nil {
		log.Printf("Saga execution failed: %v\n", err)
	} else {
		log.Println("Saga completed successfully")
	}
}
