package validation

import (
	"fmt"
	"reflect"
	"strings"
)

// IValidator GOLANG INTERFACE
// Implements a method for validating input with Validate and to get all the errors with GetErrors.
type IValidator interface {
	Validate(input interface{}, arr []string, name string)
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

func (v *Validator) Validate(input interface{}, arr []string, name string) {
	rules := NewRules()

	functionMap := map[string]any{
		"IsInt":      func(input interface{}, params string, name string) { rules.IsInt(input, name) },
		"IsString":   func(input interface{}, params string, name string) { rules.IsString(input, name) },
		"IsUUID":     func(input interface{}, params string, name string) { rules.IsUUID(input, name) },
		"IsBoolean":  func(input interface{}, params string, name string) { rules.IsBoolean(input, name) },
		"IsDatetime": func(input interface{}, params string, name string) { rules.IsDatetime(input, name) },
		"IsArray":    func(input interface{}, params string, name string) { rules.IsArray(input, name) },
		"ArrayType": func(input interface{}, params string, name string) {
			s := strings.Split(params, ":")
			rules.ArrayType(input, s[1], name)
		},
		"Length": func(input interface{}, params string, name string) {
			s := strings.Split(params, ":")
			fmt.Println(s[1])
			rules.Length(input, s[1], name)
		},
	}

	fmt.Println(arr)

	for _, value := range arr {

		functionName := value

		// Check if value contains a colon and extract the part before it
		if parts := strings.Split(value, ":"); len(parts) > 1 {
			functionName = parts[0]
		}

		if fn, exists := functionMap[functionName]; exists {
			if fn, ok := fn.(func(interface{}, string, string)); ok {
				fn(dereferenceIfNeeded(input), value, name)
			}
		}
	}

	v.errors = append(v.errors, rules.GetErrors()...)
}

func dereferenceIfNeeded(value interface{}) interface{} {
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		return reflect.ValueOf(value).Elem().Interface()
	}

	return value
}

func (v *Validator) GetErrors() []string {
	return v.errors
}

func (v *Validator) ClearErrors() {
	v.errors = []string{}
}
