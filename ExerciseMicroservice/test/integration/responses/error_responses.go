package responses

var InvalidClassIDResponseError = "Validation errors: ClassID :'56565656b' is not a valid UUID"

type ErrorType struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}
