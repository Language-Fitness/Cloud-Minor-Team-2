package helper

import "reflect"

func IsNil(input interface{}) bool {
	if input == nil {
		return true
	}

	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return true
	}

	return false
}
