package example

import (
	"gitnet.fr/deblan/go-form/form"
	"gitnet.fr/deblan/go-form/validation"
)

func CreateAddressForm() *form.Form {
	return form.NewForm(
		form.NewFieldText("Name").
			WithOptions(
				form.NewOption("label", "Name"),
				form.NewOption("required", true),
				form.NewOption("help", "A help!"),
			).
			WithConstraints(
				validation.NotBlank{},
			),
		form.NewSubForm("Address").
			WithOptions(form.NewOption("label", "Address")).
			Add(
				form.NewFieldTextarea("Street").
					WithOptions(form.NewOption("label", "Street")).
					WithConstraints(
						validation.NotBlank{},
					),
				form.NewFieldText("City").
					WithOptions(form.NewOption("label", "City")).
					WithConstraints(
						validation.NotBlank{},
					),
				form.NewFieldNumber("ZipCode").
					WithOptions(
						form.NewOption("label", "Zip code"),
						form.NewOption("help", "A field help"),
					).
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
			form.NewOption("attr", map[string]string{
				"id": "my-form",
			}),
			form.NewOption("help", "A form help!"),
		)
}
