package helper

import "reflect"

func DereferenceIfNeeded(value interface{}) interface{} {
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		return reflect.ValueOf(value).Elem().Interface()
	}

	return value
}
