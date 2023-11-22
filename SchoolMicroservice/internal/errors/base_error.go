package errors

type BaseError struct {
	Message    string
	MessageBag []string
}

func NewBaseError(message string, bag []string) *BaseError {
	return &BaseError{
		Message:    message,
		MessageBag: bag,
	}
}
