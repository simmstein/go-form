package form

import (
	"fmt"
	"time"
)

func DateBeforeMount(data any, format string) (any, error) {
	if data == nil {
		return nil, nil
	}

	switch data.(type) {
	case string:
		return data, nil
	case time.Time:
		return data.(time.Time).Format(format), nil
	case *time.Time:
		v := data.(*time.Time)
		if v != nil {
			return v.Format(format), nil
		}
	}

	return data, nil
}

// Generates an input[type=date] with default transformers
func NewFieldDate(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "date")).
		WithBeforeMount(func(data any) (any, error) {
			return DateBeforeMount(data, "2006-01-02")
		}).
		WithBeforeBind(func(data any) (any, error) {
			return time.Parse(time.DateOnly, data.(string))
		})

	return f
}

// Generates an input[type=datetime] with default transformers
func NewFieldDatetime(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "datetime")).
		WithBeforeMount(func(data any) (any, error) {
			return DateBeforeMount(data, "2006-01-02 15:04")
		}).
		WithBeforeBind(func(data any) (any, error) {
			return time.Parse("2006-01-02T15:04", data.(string))
		})

	return f
}

// Generates an input[type=datetime-local] with default transformers
func NewFieldDatetimeLocal(name string) *Field {
	f := NewField(name, "input").
		WithOptions(
			NewOption("type", "datetime-local"),
		).
		WithBeforeMount(func(data any) (any, error) {
			return DateBeforeMount(data, "2006-01-02 15:04")
		}).
		WithBeforeBind(func(data any) (any, error) {
			a, b := time.Parse("2006-01-02T15:04", data.(string))

			return a, b
		})

	return f
}

// Generates an input[type=time] with default transformers
func NewFieldTime(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "time")).
		WithBeforeMount(func(data any) (any, error) {
			return DateBeforeMount(data, "15:04")
		}).
		WithBeforeBind(func(data any) (any, error) {
			if data != nil {
				v := data.(string)

				if len(v) > 0 {
					return time.Parse(time.TimeOnly, fmt.Sprintf("%s:00", v))
				}
			}

			return nil, nil
		})

	return f
}
