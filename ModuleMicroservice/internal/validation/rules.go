package validation

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// IRules GOLANG INTERFACE
// Implements functions for validation input.
type IRules interface {
	AddError(message string)
	ClearErrors()
	GetErrors() []string
	IsInt(s interface{}) bool
	IsString(s interface{}) bool
	IsUUID(s interface{}) bool
	IsBoolean(value interface{}) bool
	IsDatetime(s interface{}) bool
	IsArray(value interface{}) bool
	ArrayType(input interface{}, expectedType string) bool
	Length(s interface{}, condition string) bool
}

// Rules GOLANG STRUCT
// Contains a list of type string errors which contains all errors.
type Rules struct {
	Errors []string
}

// NewRules GOLANG FACTORY
// Returns a Rules implementing IRules.
func NewRules() IRules {
	return &Rules{}
}

func (v *Rules) AddError(message string) {
	v.Errors = append(v.Errors, message)
}

func (v *Rules) ClearErrors() {
	v.Errors = nil
}

func (v *Rules) GetErrors() []string {
	return v.Errors
}

func (v *Rules) IsInt(s interface{}) bool {
	strValue, ok := s.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", s)
	}

	_, err := strconv.Atoi(strValue)
	if err != nil {
		v.AddError(fmt.Sprintf("'%s' is not a valid integer", strValue))
		return false
	}
	return true
}

func (v *Rules) IsString(s interface{}) bool {
	_, isString := s.(string)
	if !isString {
		v.AddError(fmt.Sprintf("'%v' is not a valid string", s))
	}
	return isString
}

func (v *Rules) IsUUID(s interface{}) bool {
	strValue, ok := s.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", s)
	}

	_, err := uuid.Parse(strValue)
	if err != nil {
		fmt.Println("test")
		v.AddError(fmt.Sprintf("'%s' is not a valid UUID", strValue))
		return false
	}
	return true
}

func (v *Rules) IsBoolean(value interface{}) bool {
	_, isBoolean := value.(bool)
	if !isBoolean {
		v.AddError(fmt.Sprintf("'%v' is not a valid boolean", value))
	}
	return isBoolean
}

func (v *Rules) IsDatetime(s interface{}) bool {
	strValue, ok := s.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", s)
	}

	_, err := time.Parse(time.RFC3339, strValue)
	if err != nil {
		v.AddError(fmt.Sprintf("'%s' is not a valid RFC3339 datetime", s))
	}
	return err == nil
}

func (v *Rules) IsArray(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Slice {
		v.AddError(fmt.Sprintf("'%v' is not a valid slice", value))
		return false
	}
	return true
}

func (v *Rules) ArrayType(input interface{}, expectedType string) bool {
	if !v.IsArray(input) {
		v.AddError("Invalid input for ArrayType")
		return false
	}

	sliceValue := reflect.ValueOf(input)

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()

		switch expectedType {
		case "string":
			if !v.IsString(element) {
				v.AddError(fmt.Sprintf("Element at index %d is not a valid string: %v", i, element))
				return false
			}
		case "int":
			if !v.IsInt(fmt.Sprintf("%v", element)) {
				v.AddError(fmt.Sprintf("Element at index %d is not a valid integer: %v", i, element))
				return false
			}
		case "bool":
			if !v.IsBoolean(element) {
				v.AddError(fmt.Sprintf("Element at index %d is not a valid boolean: %v", i, element))
				return false
			}
		case "uuid":
			if !v.IsUUID(fmt.Sprintf("%v", element)) {
				v.AddError(fmt.Sprintf("Element at index %d is not a valid UUID: %v", i, element))
				return false
			}
		default:
			v.AddError("Invalid expected type in ArrayType")
			return false
		}
	}

	return true
}

func (v *Rules) Length(s interface{}, condition string) bool {
	value, ok := s.(string)
	if !ok {
		value = fmt.Sprintf("%v", s)
	}

	operator, lengthValue := SplitCondition(condition)

	switch operator {
	case "=":
		if len(value) == lengthValue {
			return true
		} else {
			v.AddError(fmt.Sprintf("String length should be the same as %d", lengthValue))
			return false
		}
	case "<":
		if len(value) < lengthValue {
			return true
		} else {
			v.AddError(fmt.Sprintf("String length should be less than %d", lengthValue))
			return false
		}
	case "<=":
		if len(value) <= lengthValue {
			return true
		} else {
			v.AddError(fmt.Sprintf("String length should be less than or equal to %d", lengthValue))
			return false
		}
	case ">":
		if len(value) > lengthValue {
			return true
		} else {
			v.AddError(fmt.Sprintf("String length should be greater than %d", lengthValue))
			return false
		}
	case ">=":
		if len(value) >= lengthValue {
			return true
		} else {
			v.AddError(fmt.Sprintf("String length should be greater than or equal to %d", lengthValue))
			return false
		}
	default:
		v.AddError("Invalid length condition")
		return false
	}
}

func SplitCondition(condition string) (string, int) {
	re := regexp.MustCompile(`([<>]=?|<|>)?(\d+)`)
	matches := re.FindStringSubmatch(condition)
	if len(matches) != 3 {
		// Handle invalid format
		return "", -1
	}
	operator := matches[1]
	value := matches[2]
	return operator, stringToInt(value)
}

func stringToInt(s string) int {
	intValue, _ := strconv.Atoi(s)
	return intValue
}
