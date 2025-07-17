package validation

import "net/mail"

type Mail struct {
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
		errors = append(errors, Error("This value is not a valid email address."))
	}

	return errors
}
