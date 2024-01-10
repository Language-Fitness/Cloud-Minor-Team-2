package responses

var InvalidClassIDResponseError = "Validation errors: ClassID :'56565656b' is not a valid UUID"
var InvalidModuleIDResponseError = "Validation errors: ModuleID :'56565656b' is not a valid UUID"
var InvalidNameResponseError = "Validation errors: Name length should be less than 50"
var InvalidQuestionResponseError = "Validation errors: Question length should be less than 100"
var NoCorrectAnswersResponseError = "at least one answer must be correct"
var NoIncorrectAnswersResponseError = "only one answer can be correct"
var NoAnswersResponseError = "exercise must have at least two answers"

type ErrorType struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}
