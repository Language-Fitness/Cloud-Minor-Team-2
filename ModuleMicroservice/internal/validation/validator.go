package validation

import (
	"fmt"
	"strings"
)

type Validator struct {
	errors []string
}

func NewValidator() *Validator {
	return &Validator{}
}

/*
	Validate("name", ["IsString", "Length:<5"])
	Validate(9, ["IsInt"])
	Validate("["apple, "banana"]", ["IsArray", "ArrayType:string"])

	// Should map to the functions and input set values
*/

func (v *Validator) Validate(input string, arr []string) {
	rules := NewRules()

	// Create a map where keys are strings (function names) and values are function pointers
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

	for _, value := range arr {

		functionName := value // Initialize functionName with the original value

		// Check if value contains a colon and extract the part before it
		if parts := strings.Split(value, ":"); len(parts) > 1 {
			functionName = parts[0]
		}

		if fn, exists := functionMap[functionName]; exists {
			if fn, ok := fn.(func(interface{}, string)); ok {
				fn(input, value)
			}
		}
	}

	v.errors = rules.GetErrors()
}

func (v *Validator) GetErrors() []string {
	return v.errors
}
