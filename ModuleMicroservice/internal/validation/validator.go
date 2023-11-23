package validation

import (
	"fmt"
	"strings"
)

// IValidator GOLANG INTERFACE
// Implements a method for validating input with Validate and to get all the errors with GetErrors.
type IValidator interface {
	Validate(input interface{}, arr []string)
	GetErrors() []string
	ClearErrors()
}

// Validator GOLANG STRUCT
// Contains a list of type string errors which contains all errors.
type Validator struct {
	errors []string
}

// NewValidator GOLANG FACTORY
// Returns a Validator implementing IValidator.
func NewValidator() IValidator {
	fmt.Println("initilizing validator")

	return &Validator{}
}

func (v *Validator) Validate(input interface{}, arr []string) {
	rules := NewRules()

	functionMap := map[string]any{
		"IsInt":      func(input interface{}, params string) { rules.IsInt(input) },
		"IsString":   func(input interface{}, params string) { rules.IsString(input) },
		"IsUUID":     func(input interface{}, params string) { rules.IsUUID(input) },
		"IsBoolean":  func(input interface{}, params string) { rules.IsBoolean(input) },
		"IsDatetime": func(input interface{}, params string) { rules.IsDatetime(input) },
		"IsArray":    func(input interface{}, params string) { rules.IsArray(input) },
		"ArrayType": func(input interface{}, params string) {
			s := strings.Split(params, ":")
			rules.ArrayType(input, s[1])
		},
		"Length": func(input interface{}, params string) {
			s := strings.Split(params, ":")
			fmt.Println(s[1])
			rules.Length(input, s[1])
		},
	}

	fmt.Println("Input received:", input)
	fmt.Println("Input received:", arr)

	for _, value := range arr {

		functionName := value

		// Check if value contains a colon and extract the part before it
		if parts := strings.Split(value, ":"); len(parts) > 1 {
			functionName = parts[0]
		}

		if fn, exists := functionMap[functionName]; exists {
			if fn, ok := fn.(func(interface{}, string)); ok {
				fn(input, value)

				fmt.Println(rules.GetErrors())
			}
		}
	}

	v.errors = append(rules.GetErrors(), v.errors...)
	fmt.Println(v.errors)
}

func (v *Validator) GetErrors() []string {
	return v.errors
}

func (v *Validator) ClearErrors() {
	v.errors = []string{}
}
