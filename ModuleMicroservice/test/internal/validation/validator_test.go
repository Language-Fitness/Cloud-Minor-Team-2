package validation

import (
	"Module/internal/validation"
	"fmt"
	_ "fmt"
	_ "reflect"
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	// Initialize your Validator (v) and Rules (rules) as needed

	// Define test cases
	testCases := []struct {
		name  string
		input string
		arr   []string
	}{
		{
			name:  "test IsInt",
			input: "42",
			arr:   []string{"IsInt"},
		},
		{
			name:  "test IsString",
			input: "Hello, World!",
			arr:   []string{"IsString", "Length:<5"},
		},
		// Add more test cases for other validation functions
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Initialize your Validator (v) and Rules (rules) as needed for each test case
			v := validation.NewValidator()

			// Call the Validate function
			v.Validate(tc.input, tc.arr, "string")

			// Add assertions here to check if the validation was successful

			fmt.Println(v.GetErrors())
		})
	}

}
