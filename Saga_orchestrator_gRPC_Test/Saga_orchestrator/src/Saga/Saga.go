package Saga

import (
	"fmt"
	"sync"
)

var sagaStore = make(map[string]SagaInstance)
var mu sync.Mutex

func CreateSaga(steps ...SagaStep) string {
	mu.Lock()
	defer mu.Unlock()

	sagaID := fmt.Sprintf("saga-%d", len(sagaStore)+1)
	saga := SagaInstance{
		ID:          sagaID,
		CurrentStep: 0,
		Steps:       steps,
	}

	sagaStore[sagaID] = saga

	return sagaID
}

func ExecuteSaga(sagaID string) error {
	mu.Lock()
	defer mu.Unlock()

	saga, exists := sagaStore[sagaID]
	if !exists {
		return fmt.Errorf("Saga not found")
	}

	for i := saga.CurrentStep; i < len(saga.Steps); i++ {
		step := saga.Steps[i]
		if err := step.Execute(); err != nil {
			// Handle step failure by initiating compensating actions
			for j := i; j >= 0; j-- {
				compensateErr := step.Compensate()
				if compensateErr != nil {
					return fmt.Errorf("Saga compensation error: %s", compensateErr)
				}
			}
			return fmt.Errorf("Saga step failed: %s", err)
		}
		saga.CurrentStep = i + 1
	}

	// Saga completed successfully
	delete(sagaStore, sagaID)
	return nil
}
