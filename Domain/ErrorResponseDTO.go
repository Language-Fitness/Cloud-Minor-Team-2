package Domain

type ErrorResponseDTO struct {
	StatusText string
	StatusCode int
	ErrorType  string
	Error      string
}
