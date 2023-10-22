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
	// Define any fields or rules you need for validation.
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func (v *Validator) IsString(s interface{}) bool {
	_, isString := s.(string)
	return isString
}

func (v *Validator) IsUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

func (v *Validator) IsBoolean(value interface{}) bool {
	_, isBoolean := value.(bool)
	return isBoolean
}

func (v *Validator) IsDatetime(s string) bool {
	_, err := time.Parse(time.RFC3339, s)
	return err == nil
}

func (v *Validator) IsArray(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	return valueType.Kind() == reflect.Slice
}

func (v *Validator) ArrayType(input interface{}, expectedType string) bool {
	if !v.IsArray(input) {
		return false
	}

	sliceValue := reflect.ValueOf(input)

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()

		// Perform type-specific validation based on expectedType
		switch expectedType {
		case "string":
			if !v.IsString(element) {
				return false
			}
		case "int":
			if !v.IsInt(fmt.Sprintf("%v", element)) {
				return false
			}
		case "bool":
			if !v.IsBoolean(element) {
				return false
			}
		case "uuid":
			if !v.IsUUID(fmt.Sprintf("%v", element)) {
				return false
			}
		default:
			// Handle other expected types or return an error if needed
			return false
		}
	}

	return true
}

func (v *Validator) Length(value string, condition string) bool {
	operator, lengthValue := SplitCondition(condition)
	fmt.Printf("Operator: %s, Value: %s\n", operator, value)
	fmt.Printf(string(rune(len(value))), lengthValue)

	switch operator {
	case "<":
		return len(value) < lengthValue
	case "<=":
		return len(value) <= lengthValue
	case ">":
		return len(value) > lengthValue
	case ">=":
		return len(value) >= lengthValue
	default:
		// Handle invalid operator
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
