package validation

import "net/mail"

type Mail struct {
	Message string
}

func NewMail() Mail {
	return Mail{
		Message: "This value is not a valid email address.",
	}
}

func (c Mail) Validate(data any) []Error {
	errors := []Error{}

	notBlank := NotBlank{}
	nbErrs := notBlank.Validate(data)

	if len(nbErrs) > 0 {
		return errors
	}

	_, err := mail.ParseAddress(data.(string))

	if err != nil {
		errors = append(errors, Error(c.Message))
	}

	return errors
}
