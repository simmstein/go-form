package example

import (
	"gitnet.fr/deblan/go-form/form"
	"gitnet.fr/deblan/go-form/validation"
)

type Tag struct {
	Name string
}

type Post struct {
	Tags  []Tag
	Tags2 []Tag
	Tags3 Tag
	Tag   Tag
}

func CreateExampleForm2() *form.Form {
	tags := []Tag{Tag{"tag1"}, Tag{"tag2"}, Tag{"tag3"}}

	choices := form.NewChoices(tags).
		WithLabelBuilder(func(key int, item any) string {
			return item.(Tag).Name
		})

	return form.NewForm(
		form.NewFieldChoice("Tag").
			WithOptions(
				form.NewOption("choices", choices),
				form.NewOption("label", "Tag"),
				form.NewOption("required", true),
			).
			WithConstraints(
				validation.NotBlank{},
			),
		form.NewFieldChoice("Tags").
			WithSlice().
			WithOptions(
				form.NewOption("choices", choices),
				form.NewOption("label", "Tags"),
				form.NewOption("multiple", true),
				form.NewOption("required", true),
			).
			WithConstraints(
				validation.NotBlank{},
			),
		form.NewFieldChoice("Tags2").
			WithSlice().
			WithOptions(
				form.NewOption("choices", choices),
				form.NewOption("label", "Tags"),
				form.NewOption("multiple", true),
				form.NewOption("expanded", true),
				form.NewOption("required", true),
			).
			WithConstraints(
				validation.NotBlank{},
			),
		form.NewFieldChoice("Tag3").
			WithOptions(
				form.NewOption("choices", choices),
				form.NewOption("label", "Tag"),
				form.NewOption("expanded", true),
			),
		form.NewSubmit("submit"),
	).
		End().
		WithMethod("POST").
		WithAction("/")
}

func CreateExampleForm() *form.Form {
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
		form.NewFieldDate("Date").WithOptions(form.NewOption("label", "Date")),
		// form.NewFieldDatetime("DateTime").WithOptions(form.NewOption("label", "DateTime")),
		form.NewFieldDatetimeLocal("DateTime").WithOptions(form.NewOption("label", "DateTimeLocal")),
		form.NewFieldTime("Time").WithOptions(form.NewOption("label", "Time")),
		form.NewFieldCheckbox("Checkbox").WithOptions(form.NewOption("label", "Checkbox")),
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
				form.NewFieldRange("Foo").
					WithOptions(
						form.NewOption("label", "Foo"),
					),
				form.NewFieldMail("Email").
					WithOptions(
						form.NewOption("label", "Email"),
					).
					WithConstraints(
						validation.NotBlank{},
						validation.Mail{},
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
