package custom_erros

type ValidationError struct {
	BaseError
}

func NewValidationError(bag []string) *ValidationError {
	return &ValidationError{
		BaseError: BaseError{
			Message:    "Validation Errors:",
			MessageBag: bag,
		},
	}
}
