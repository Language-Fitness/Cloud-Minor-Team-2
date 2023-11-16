package errors

type AuthError struct {
	BaseError
}

func NewAuthError(bag []string) *AuthError {
	return &AuthError{
		BaseError: BaseError{
			Message:    "Auth Errors:",
			MessageBag: bag,
		},
	}
}
