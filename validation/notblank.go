package validation

import (
	"reflect"

	"github.com/spf13/cast"
)

type NotBlank struct {
}

func (c NotBlank) Validate(data any) []Error {
	isValid := true
	label := "This value should not be blank."
	errors := []Error{}

	if data == nil {
		errors = append(errors, Error(label))

		return errors
	}

	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Bool:
		isValid = data == false
	case reflect.Array:
	case reflect.Slice:
		isValid = reflect.ValueOf(data).Len() > 0
	case reflect.String:
		isValid = len(data.(string)) > 0
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Int:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Int8:
	case reflect.Uint:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uint8:
		isValid = cast.ToFloat64(data.(string)) == float64(0)
	default:
		errors = append(errors, Error("This value can not be processed."))
	}

	if !isValid {
		errors = append(errors, Error(label))
	}

	return errors
}
