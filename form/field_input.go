package form

import (
	"github.com/spf13/cast"
)

// Generates an input[type=text]
func NewFieldText(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "text"))

	return f
}

// Generates an input[type=number] with default transformers
func NewFieldNumber(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "number")).
		WithBeforeBind(func(data any) (any, error) {
			return cast.ToFloat64(data), nil
		})

	return f
}

// Generates an input[type=email]
func NewFieldMail(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "email"))

	return f
}

// Generates an input[type=range]
func NewFieldRange(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "range")).
		WithBeforeBind(func(data any) (any, error) {
			return cast.ToFloat64(data), nil
		})

	return f
}

// Generates an input[type=password]
func NewFieldPassword(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "password"))

	return f
}

// Generates an input[type=hidden]
func NewFieldHidden(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "hidden"))

	return f
}

// Generates an input[type=submit]
func NewSubmit(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "submit"))

	f.Data = "Submit"

	return f
}
