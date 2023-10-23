package validation

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

type Validator struct {
	Errors []string
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) AddError(message string) {
	v.Errors = append(v.Errors, message)
}

func (v *Validator) ClearErrors() {
	v.Errors = nil
}

func (v *Validator) GetErrors() []string {
	return v.Errors
}

func (v *Validator) IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(fmt.Sprintf("'%s' is not a valid integer", s))
	}
	return err == nil
}

func (v *Validator) IsString(s interface{}) bool {
	_, isString := s.(string)
	if !isString {
		v.AddError(fmt.Sprintf("'%v' is not a valid string", s))
	}
	return isString
}

func (v *Validator) IsUUID(s string) bool {
	_, err := uuid.Parse(s)
	if err != nil {
		v.AddError(fmt.Sprintf("'%s' is not a valid UUID", s))
	}
	return err == nil
}

func (v *Validator) IsBoolean(value interface{}) bool {
	_, isBoolean := value.(bool)
	if !isBoolean {
		v.AddError(fmt.Sprintf("'%v' is not a valid boolean", value))
	}
	return isBoolean
}

func (v *Validator) IsDatetime(s string) bool {
	_, err := time.Parse(time.RFC3339, s)
	if err != nil {
		v.AddError(fmt.Sprintf("'%s' is not a valid RFC3339 datetime", s))
	}
	return err == nil
}

func (v *Validator) IsArray(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Slice {
		v.AddError(fmt.Sprintf("'%v' is not a valid slice", value))
		return false
	}
	return true
}

func (v *Validator) ArrayType(input interface{}, expectedType string) bool {
	if !v.IsArray(input) {
		v.AddError("Invalid input for ArrayType")
		return false
	}

	sliceValue := reflect.ValueOf(input)

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()

		// Perform type-specific validation based on expectedType
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

func (v *Validator) Length(value string, condition string) bool {
	operator, lengthValue := SplitCondition(condition)

	switch operator {
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
		// Handle invalid operator
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
