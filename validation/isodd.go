package validation

import (
	"strconv"
)

type IsOdd struct {
	Message          string
	TypeErrorMessage string
}

func NewIsOdd() IsOdd {
	return IsOdd{
		Message:          "This value is not a odd number.",
		TypeErrorMessage: "This value can not be processed.",
	}
}

func (c IsOdd) Validate(data any) []Error {
	errors := []Error{}

	// The constraint should not validate an empty data
	if len(NewNotBlank().Validate(data)) == 0 {
		i, err := strconv.Atoi(data.(string))

		if err != nil {
			errors = append(errors, Error(c.TypeErrorMessage))
		} else if i%2 != 1 {
			errors = append(errors, Error(c.Message))
		}
	}

	return errors
}
