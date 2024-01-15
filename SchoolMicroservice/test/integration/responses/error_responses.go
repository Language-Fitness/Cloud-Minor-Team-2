package responses

var SchoolNotFoundResponseError = "school not found"
var InvalidPermissionResponseError = "invalid permissions for this action"

type ErrorType struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}
