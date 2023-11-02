package Saga

import (
	"Saga_orchestrator/src/Repository"
	"Saga_orchestrator/src/Steps"
	"Saga_orchestrator/src/models"
	"github.com/google/uuid"
	"sync"
)

var mu sync.Mutex

func CreateSaga(steps ...models.SagaStep) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	// Create a new UUID for the Saga ID
	sagaID := uuid.New().String()

	// Create the SagaInstance with an initial state of "NotStarted"
	newSaga := &models.SagaInstance{
		ID:          sagaID,
		CurrentStep: 0,
		State:       models.NotStarted,
		Steps:       steps,
	}

	if err := Repository.CreateSagaInMongoDB(newSaga); err != nil {
		return "", err
	}

	return sagaID, nil
}

func ExecuteSaga(sagaID string) error {
	mu.Lock()
	defer mu.Unlock()

	// Retrieve the saga state from MongoDB
	saga, err := Repository.GetSagaFromMongoDB(sagaID)
	if err != nil {
		return err
	}

	// Update the state to "InProgress"
	saga.State = models.InProgress
	if err := Repository.UpdateSagaInMongoDB(saga); err != nil {
		return err
	}

	compensate := false // Track whether compensation is needed

	for i := saga.CurrentStep; i < len(saga.Steps); i++ {
		step := saga.Steps[i]
		Execute := Steps.MapExecuteFuncName(step)
		if Execute != nil {
			if err := Execute(step.ExecuteParams); err != nil {
				// Handle step failure
				saga.State = models.Failed
				compensate = true // Set the compensation flag
				break             // Exit the loop on step failure
			}

			saga.CurrentStep = i + 1
			// If all steps are successfully completed, set the state to "Completed"
			if saga.CurrentStep == len(saga.Steps) {
				saga.State = models.Completed
			}

			// Update the saga state in MongoDB after each step
			if err := Repository.UpdateSagaInMongoDB(saga); err != nil {
				return err
			}
		}
	}

	if compensate {
		// If compensation is needed, execute compensation steps
		for j := saga.CurrentStep - 1; j >= 0; j-- {
			step := saga.Steps[j]
			Compensation := Steps.MapCompensationFuncName(step)
			if Compensation != nil {
				if err := Compensation(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
