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
	IsInt(s interface{}, name string) bool
	IsString(s interface{}, name string) bool
	IsUUID(s interface{}, name string) bool
	IsBoolean(value interface{}, name string) bool
	IsDatetime(s interface{}, name string) bool
	IsArray(value interface{}, name string) bool
	ArrayType(input interface{}, expectedType string, name string) bool
	Length(s interface{}, condition string, name string) bool
	Size(s interface{}, condition string, name string) bool
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

func (v *Rules) IsInt(s interface{}, name string) bool {
	strValue, ok := s.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", s)
	}

	_, err := strconv.Atoi(strValue)
	if err != nil {
		v.AddError(fmt.Sprintf("%s: '%v' is not a valid integer", name, strValue))
		return false
	}
	return true
}

func (v *Rules) IsString(s interface{}, name string) bool {
	_, isString := s.(string)
	if !isString {
		v.AddError(fmt.Sprintf("%s: '%v' is not a valid string", name, s))
	}
	return isString
}

func (v *Rules) IsUUID(s interface{}, name string) bool {
	strValue, ok := s.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", s)
	}

	_, err := uuid.Parse(strValue)
	if err != nil {
		v.AddError(fmt.Sprintf("%s :'%v' is not a valid UUID", name, strValue))
		return false
	}
	return true
}

func (v *Rules) IsBoolean(value interface{}, name string) bool {
	_, isBoolean := value.(bool)
	if !isBoolean {
		v.AddError(fmt.Sprintf("%s : '%v' is not a valid boolean", name, value))
	}
	return isBoolean
}

func (v *Rules) IsDatetime(s interface{}, name string) bool {
	strValue, ok := s.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", s)
	}

	_, err := time.Parse(time.RFC3339, strValue)
	if err != nil {
		v.AddError(fmt.Sprintf("%s : '%v' is not a valid RFC3339 datetime", name, s))
	}
	return err == nil
}

func (v *Rules) IsArray(value interface{}, name string) bool {
	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Slice {
		v.AddError(fmt.Sprintf("%s : '%v' is not a valid slice", name, value))
		return false
	}
	return true
}

func (v *Rules) ArrayType(input interface{}, expectedType string, name string) bool {
	if !v.IsArray(input, name) {
		v.AddError("Invalid input for ArrayType")
		return false
	}

	sliceValue := reflect.ValueOf(input)

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()

		switch expectedType {
		case "string":
			if !v.IsString(element, name) {
				v.AddError(fmt.Sprintf("%s at index %d is not a valid string: %v", name, i, element))
				return false
			}
		case "int":
			if !v.IsInt(fmt.Sprintf("%v", element), name) {
				v.AddError(fmt.Sprintf("%s at index %d is not a valid integer: %v", name, i, element))
				return false
			}
		case "bool":
			if !v.IsBoolean(element, name) {
				v.AddError(fmt.Sprintf("%s at index %d is not a valid boolean: %v", name, i, element))
				return false
			}
		case "uuid":
			if !v.IsUUID(fmt.Sprintf("%v", element), name) {
				v.AddError(fmt.Sprintf("%s at index %d is not a valid UUID: %v", name, i, element))
				return false
			}
		default:
			v.AddError("Invalid expected type in ArrayType")
			return false
		}
	}

	return true
}

func (v *Rules) Length(s interface{}, condition string, name string) bool {
	value, ok := s.(string)
	if !ok {
		value = fmt.Sprintf("%v", s)
	}

	operator, lengthValue := SplitCondition(condition)

	switch operator {
	case "=":
		if len(value) == lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s should have a length equal to %d", name, lengthValue))
	case "<":
		if len(value) < lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s length should be less than %d", name, lengthValue))
	case "<=":
		if len(value) <= lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s length should be less than or equal to %d", name, lengthValue))
	case ">":
		if len(value) > lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s length should be greater than %d", name, lengthValue))
	case ">=":
		if len(value) >= lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s length should be greater than or equal to %d", name, lengthValue))
	default:
		v.AddError("Invalid length condition")
	}

	return false
}

func (v *Rules) Size(s interface{}, condition string, name string) bool {
	strValue, ok := s.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", s)
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		v.AddError(fmt.Sprintf("%s: '%v' is not a valid integer", name, strValue))
		return false
	}

	operator, lengthValue := SplitCondition(condition)

	switch operator {
	case "=":
		if value == lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s should have a size equal to %d", name, lengthValue))
	case "<":
		if value < lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s size should be less than %d", name, lengthValue))
	case "<=":
		if value <= lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s size should be less than or equal to %d", name, lengthValue))
	case ">":
		if value > lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s size should be greater than %d", name, lengthValue))
	case ">=":
		if value >= lengthValue {
			return true
		}
		v.AddError(fmt.Sprintf("%s size should be greater than or equal to %d", name, lengthValue))
	default:
		v.AddError("Invalid length condition")
	}

	return false
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
