package validation

import (
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

type Range struct {
	Min              *float64
	Max              *float64
	MinMessage       string
	MaxMessage       string
	RangeMessage     string
	TypeErrorMessage string
}

func NewRange() Range {
	return Range{
		MinMessage:       "This value must be greater than or equal to {{ min }}.",
		MaxMessage:       "This value must be less than or equal to {{ max }}.",
		RangeMessage:     "This value should be between {{ min }} and {{ max }}.",
		TypeErrorMessage: "This value can not be processed.",
	}
}

func (c Range) WithMin(v float64) Range {
	c.Min = &v

	return c
}

func (c Range) WithMax(v float64) Range {
	c.Max = &v

	return c
}

func (c Range) WithRange(vMin, vMax float64) Range {
	c.Min = &vMin
	c.Max = &vMax

	return c
}

func (c Range) Validate(data any) []Error {
	if c.Min == nil && c.Max == nil {
		return []Error{}
	}

	errors := []Error{}

	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
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
	case reflect.String:
		isValidMin := c.Min == nil || *c.Min <= cast.ToFloat64(data.(string))
		isValidMax := c.Max == nil || *c.Max >= cast.ToFloat64(data.(string))

		if !isValidMin || !isValidMax {
			errors = append(errors, Error(c.BuildMessage()))
		}
	default:
		errors = append(errors, Error(c.TypeErrorMessage))
	}

	return errors
}

func (c *Range) BuildMessage() string {
	var message string

	if c.Min != nil && c.Max == nil {
		message = c.MinMessage
	} else if c.Max != nil && c.Min == nil {
		message = c.MaxMessage
	} else {
		message = c.RangeMessage
	}

	message = strings.ReplaceAll(message, "{{ min }}", cast.ToString(c.Min))
	message = strings.ReplaceAll(message, "{{ max }}", cast.ToString(c.Max))

	return message
}
