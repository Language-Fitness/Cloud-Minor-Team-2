package helper

import (
	"errors"
	"fmt"
)

func AggregateErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	errorMessages := make([]string, len(errs))
	for i, err := range errs {
		errorMessages[i] = err.Error()
	}

	aggregateErrorMessage := fmt.Sprintf("Multiple errors: %v", errorMessages)
	return errors.New(aggregateErrorMessage)
}
