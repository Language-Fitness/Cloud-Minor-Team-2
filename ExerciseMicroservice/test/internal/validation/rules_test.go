package validation

import (
	"ExerciseMicroservice/internal/validation"
	"fmt"
	"reflect"
	"testing"
)

func TestRules_IsInt(t *testing.T) {
	validator := validation.NewRules()

	// Test cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"-456", true},
		{"abc", false},
		{"123.45", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := validator.IsInt(tc.input, "Test")
			if result != tc.expected {
				t.Errorf("IsInt(%s) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestRules_IsString(t *testing.T) {
	validator := validation.NewRules()

	// Test cases
	testCases := []struct {
		input    interface{}
		expected bool
	}{
		{"This is a string", true},
		{123, false},
		{true, false},
		{3.14, false},
	}

	for _, tc := range testCases {
		t.Run("IsString with input "+toString(tc.input), func(t *testing.T) {
			result := validator.IsString(tc.input, "Test")
			if result != tc.expected {
				t.Errorf("IsString(%v) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func toString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

func TestRules_IsUUID(t *testing.T) {
	validator := validation.NewRules()

	// Test cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{"6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},
		{"not-a-uuid", false},
		{"00000000-0000-0000-0000-000000000000", true},
		{"6ba7b810-9dad-11d1-80b4-00c04fd430c8Z", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := validator.IsUUID(tc.input, "Test")
			if result != tc.expected {
				t.Errorf("IsUUID(%s) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestRules_IsBoolean(t *testing.T) {
	validator := validation.NewRules()

	// Test cases
	testCases := []struct {
		input    interface{}
		expected bool
	}{
		{true, true},
		{false, true},
		{"not-a-boolean", false},
		{42, false},
	}

	for _, tc := range testCases {
		t.Run("IsBoolean with input "+toString(tc.input), func(t *testing.T) {
			result := validator.IsBoolean(tc.input, "Test")
			if result != tc.expected {
				t.Errorf("IsBoolean(%v) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestRules_IsDatetime(t *testing.T) {
	validator := validation.NewRules()

	// Test cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{"2023-10-22T12:00:00Z", true},
		{"not-a-datetime", false},
		{"2023-10-22", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := validator.IsDatetime(tc.input, "Test")
			if result != tc.expected {
				t.Errorf("IsDatetime(%s) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestRules_IsArray(t *testing.T) {
	validator := validation.NewRules()

	// Test cases
	testCases := []struct {
		input    interface{}
		expected bool
	}{
		{[]int{1, 2, 3}, true},
		{[]string{"apple", "banana", "cherry"}, true},
		{"not-an-array", false},
		{42, false},
	}

	for _, tc := range testCases {
		t.Run(reflect.TypeOf(tc.input).String(), func(t *testing.T) {
			result := validator.IsArray(tc.input, "Test")
			if result != tc.expected {
				t.Errorf("IsArray(%v) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestRules_ArrayType(t *testing.T) {
	validator := validation.NewRules()

	// Test cases
	testCases := []struct {
		input        interface{}
		expectedType string
		expected     bool
	}{
		{[]int{1, 2, 3}, "int", true},
		{[]string{"apple", "banana", "cherry"}, "string", true},
		{[]string{"apple", "banana", "cherry"}, "int", false},
		{[]bool{true, false}, "bool", true},
		{[]bool{true, false}, "uuid", false},
		{"not-an-array", "string", false},
		{42, "int", false},
	}

	for _, tc := range testCases {
		t.Run(reflect.TypeOf(tc.input).String()+"_"+tc.expectedType, func(t *testing.T) {
			result := validator.ArrayType(tc.input, tc.expectedType, "Test")
			if result != tc.expected {
				t.Errorf("ArrayType(%v, %s) = %v; expected %v", tc.input, tc.expectedType, result, tc.expected)
			}
		})
	}
}

func TestRules_Length(t *testing.T) {
	validator := validation.NewRules()

	testCases := []struct {
		value     string
		condition string
		expected  bool
	}{
		{"example", "<=7", true},
		{"example", "<=17", true},
		{"example", "<=3", false},
		{"example", ">=5", true},
		{"example", ">=3", true},
		{"example", ">=17", false},
		{"example", "<8", true},
		{"example", "<4", false},
		{"example", "<18", true},
		{"example", ">4", true},
		{"example", ">2", true},
		{"example", ">18", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Length(%s, %s)", tc.value, tc.condition), func(t *testing.T) {
			result := validator.Length(tc.value, tc.condition, "Test")
			if result != tc.expected {
				t.Errorf("Length(%s, %s) = %v; expected %v", tc.value, tc.condition, result, tc.expected)
			}
		})
	}
}
