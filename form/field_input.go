package form

import (
	"github.com/spf13/cast"
)

func NewFieldText(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "text"))

	return f
}

func NewFieldNumber(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "number")).
		WithBeforeBind(func(data any) (any, error) {
			return cast.ToFloat64(data), nil
		})

	return f
}

func NewFieldMail(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "email"))

	return f
}

func NewFieldRange(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "range")).
		WithBeforeBind(func(data any) (any, error) {
			return cast.ToFloat64(data), nil
		})

	return f
}

func NewFieldPassword(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "password"))

	return f
}

func NewSubmit(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "submit"))

	f.Data = "Submit"

	return f
}
