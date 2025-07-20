package form

import (
	"github.com/spf13/cast"
)

// Generates an input[type=checkbox]
func NewFieldCheckbox(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "checkbox")).
		WithBeforeMount(func(data any) (any, error) {
			switch data.(type) {
			case string:
				data = data == "1"
			case bool:
				return data, nil
			}

			return cast.ToInt(data), nil
		}).
		WithBeforeBind(func(data any) (any, error) {
			switch data.(type) {
			case string:
				return data == "1", nil
			case bool:
				return data, nil
			}

			return cast.ToBool(data), nil
		})

	return f
}
