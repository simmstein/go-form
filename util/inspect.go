package util

import (
	"errors"
	"reflect"
)

func InspectStruct(input interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, errors.New("Invalid type")
	}

	result := make(map[string]interface{})

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		result[field.Name] = value.Interface()
	}

	return result, nil
}
