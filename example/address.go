package example

import (
	"gitnet.fr/deblan/go-form/form"
	"gitnet.fr/deblan/go-form/validation"
)

func CreateAddressForm() *form.Form {
	return form.NewForm(
		form.NewFieldText("Name").
			WithOptions(
				form.Option{Name: "label", Value: "Name"},
				form.Option{Name: "required", Value: true},
			).
			WithConstraints(
				validation.NotBlank{},
			),
		form.NewSubForm("Address").
			WithOptions(form.Option{Name: "label", Value: "Address"}).
			Add(
				form.NewFieldTextarea("Street").
					WithOptions(form.Option{Name: "label", Value: "Street"}).
					WithConstraints(
						validation.NotBlank{},
					),
				form.NewFieldText("City").
					WithOptions(form.Option{Name: "label", Value: "City"}).
					WithConstraints(
						validation.NotBlank{},
					),
				form.NewFieldNumber("ZipCode").
					WithOptions(form.Option{Name: "label", Value: "Zip code"}).
					WithConstraints(
						validation.NotBlank{},
					),
			),
		form.NewSubmit("submit"),
	).
		End().
		WithMethod("POST").
		// WithMethod("GET").
		WithAction("/").
		WithOptions(
			form.Option{Name: "attr", Value: map[string]string{
				"id": "my-form",
			}},
		)
}
