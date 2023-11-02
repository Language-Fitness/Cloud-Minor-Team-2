package Steps

import (
	"Saga_orchestrator/src/Steps/AddressSteps"
	"Saga_orchestrator/src/Steps/UserSteps"
	"Saga_orchestrator/src/models"
)

func MapExecuteFuncName(step models.SagaStep) func([]string) error {
	executeFunction, exists := ExecuteFunctions[step.Execute]
	if !exists {
		return nil
	}
	return executeFunction
}

func MapCompensationFuncName(step models.SagaStep) func() error {
	compensateFunction, exists := CompensateFunctions[step.Compensate]
	if !exists {
		return nil
	}
	return compensateFunction
}

var ExecuteFunctions = map[string]func([]string) error{
	UserSteps.UserStep1Execute:       UserSteps.Step1Execute,
	AddressSteps.AddressStep1Execute: AddressSteps.Step1Execute,
}

var CompensateFunctions = map[string]func() error{
	UserSteps.UserStep1Compensate:       UserSteps.Step1Compensate,
	AddressSteps.AddressStep1Compensate: AddressSteps.Step1Compensate,
}
