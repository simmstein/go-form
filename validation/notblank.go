package validation

import (
	"fmt"
	"reflect"
)

type NotBlank struct {
}

func (c NotBlank) Validate(data any) []Error {
	isValid := true
	errors := []Error{}

	fmt.Printf("%+v\n", data)

	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if data == nil {
		isValid = false
	} else if t.Kind() == reflect.Bool {
		if data == false {
			isValid = false
		}
	} else if t.Kind() == reflect.Array {
		if len(data.([]interface{})) == 0 {
			isValid = false
		}
	} else if t.Kind() == reflect.String {
		if len(data.(string)) == 0 {
			isValid = false
		}
	} else {
		fmt.Printf("d=%+v\n", data)
	}

	if !isValid {
		errors = append(errors, Error("This value should be blank"))
	}

	return errors
}
