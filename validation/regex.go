package validation

import (
	"reflect"
	"regexp"
)

type Regex struct {
	Message          string
	TypeErrorMessage string
	Match            bool
	Expression       string
}

func NewRegex(expr string) Regex {
	return Regex{
		Message:          "This value is not valid.",
		TypeErrorMessage: "This value can not be processed.",
		Match:            true,
		Expression:       expr,
	}
}

func (c Regex) MustMatch() Regex {
	c.Match = true

	return c
}

func (c Regex) MustNotMatch() Regex {
	c.Match = false

	return c
}

func (c Regex) Validate(data any) []Error {
	errors := []Error{}
	notBlank := NotBlank{}
	nbErrs := notBlank.Validate(data)

	if len(nbErrs) > 0 {
		return errors
	}

	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.String:
		matched, _ := regexp.MatchString(c.Expression, data.(string))

		if !matched && c.Match || matched && !c.Match {
			errors = append(errors, Error(c.Message))
		}

	default:
		errors = append(errors, Error(c.TypeErrorMessage))
	}

	return errors
}
