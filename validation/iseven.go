package validation

import (
	"strconv"
)

type IsEven struct {
	Message          string
	TypeErrorMessage string
}

func NewIsEven() IsEven {
	return IsEven{
		Message:          "This value is not an even number.",
		TypeErrorMessage: "This value can not be processed.",
	}
}

func (c IsEven) Validate(data any) []Error {
	errors := []Error{}

	// The constraint should not validate an empty data
	if len(NewNotBlank().Validate(data)) == 0 {
		i, err := strconv.Atoi(data.(string))

		if err != nil {
			errors = append(errors, Error(c.TypeErrorMessage))
		} else if i%2 != 0 {
			errors = append(errors, Error(c.Message))
		}
	}

	return errors
}
