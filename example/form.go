package example

import (
	"net/http"
	"time"

	"github.com/spf13/cast"
	"gitnet.fr/deblan/go-form/form"
	"gitnet.fr/deblan/go-form/validation"
)

type Item struct {
	Id   uint
	Name string
}

type ExampleOtherInputs struct {
	Number   float32
	Range    uint
	Mail     string
	Password string
}

type ExampleChoices struct {
	Select                 *Item
	SelectExpanded         *Item
	MultipleSelect         []Item
	MultipleSelectExpanded []Item
}

type ExampleDates struct {
	Date          *time.Time
	DateTime      *time.Time
	DateTimeLocal *time.Time
	Time          *time.Time
}

type ExampleData struct {
	Bytes    []byte
	Text     string
	Checkbox bool
	Dates    ExampleDates
	Choices  ExampleChoices
	Inputs   ExampleOtherInputs
}

func CreateDataForm() *form.Form {
	items := []Item{
		Item{Id: 1, Name: "Item 1"},
		Item{Id: 2, Name: "Item 2"},
		Item{Id: 3, Name: "Item 3"},
	}

	itemsChoices := form.NewChoices(items).
		WithValueBuilder(func(key int, item any) string {
			return cast.ToString(item.(Item).Id)
		}).
		WithLabelBuilder(func(key int, item any) string {
			return item.(Item).Name
		})

	return form.NewForm(
		form.NewFieldText("Bytes").
			WithOptions(
				form.NewOption("label", "Bytes"),
				form.NewOption("required", true),
				form.NewOption("row_attr", form.Attrs{
					"class": "col-12 mb-3",
				}),
			).
			WithBeforeMount(func(data any) (any, error) {
				return cast.ToString(data), nil
			}).
			WithBeforeBind(func(data any) (any, error) {
				return []byte(cast.ToString(data)), nil
			}).
			WithConstraints(
				validation.NewNotBlank(),
			),
		form.NewFieldTextarea("Text").
			WithOptions(
				form.NewOption("label", "Text"),
				form.NewOption("help", "Must contain 'deblan'"),
				form.NewOption("row_attr", form.Attrs{
					"class": "col-12 mb-3",
				}),
			).
			WithConstraints(
				validation.NewRegex(`deblan`),
			),
		form.NewFieldCheckbox("Checkbox").
			WithOptions(
				form.NewOption("label", "Checkbox"),
				form.NewOption("row_attr", form.Attrs{
					"class": "col-12 mb-3",
				}),
			),
		form.NewSubForm("Inputs").
			WithOptions(
				form.NewOption("label", "Inputs"),
			).
			Add(
				form.NewFieldNumber("Number").
					WithOptions(
						form.NewOption("label", "Number"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					).
					WithConstraints(
						validation.NewRange().WithRange(1, 20),
						validation.NewIsEven(),
					),
				form.NewFieldRange("Range").
					WithOptions(
						form.NewOption("label", "Range"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					),
				form.NewFieldMail("Mail").
					WithOptions(
						form.NewOption("label", "Mail"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					).
					WithConstraints(
						validation.Mail{},
					),
				form.NewFieldPassword("Password").
					WithOptions(
						form.NewOption("label", "Password"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					).
					WithConstraints(
						validation.NewLength().WithMin(10),
					),
			),
		form.NewSubForm("Dates").
			WithOptions(
				form.NewOption("label", "Dates"),
			).
			Add(
				form.NewFieldDate("Date").
					WithOptions(
						form.NewOption("label", "Date"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					),
				form.NewFieldDatetime("DateTime").
					WithOptions(
						form.NewOption("label", "Datetime"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					),
				form.NewFieldDatetimeLocal("DateTimeLocal").
					WithOptions(
						form.NewOption("label", "DateTime local"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					),
				form.NewFieldTime("Time").
					WithOptions(
						form.NewOption("label", "Time"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					),
			),
		form.NewSubForm("Choices").
			WithOptions(form.NewOption("label", "Choices")).
			Add(
				form.NewFieldChoice("Select").
					WithOptions(
						form.NewOption("choices", itemsChoices),
						form.NewOption("label", "Select"),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					).
					WithConstraints(
						validation.NewNotBlank(),
					),
				form.NewFieldChoice("SelectExpanded").
					WithOptions(
						form.NewOption("choices", itemsChoices),
						form.NewOption("label", "Select (expanded)"),
						form.NewOption("expanded", true),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					),
				form.NewFieldChoice("MultipleSelect").
					WithSlice().
					WithOptions(
						form.NewOption("choices", itemsChoices),
						form.NewOption("label", "Multiple select"),
						form.NewOption("multiple", true),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					).
					WithConstraints(
						validation.NewNotBlank(),
						validation.NewLength().WithExact(2),
					),
				form.NewFieldChoice("MultipleSelectExpanded").
					WithSlice().
					WithOptions(
						form.NewOption("choices", itemsChoices),
						form.NewOption("label", "Multiple select (expanded)"),
						form.NewOption("expanded", true),
						form.NewOption("multiple", true),
						form.NewOption("row_attr", form.Attrs{
							"class": "col-12 mb-3",
						}),
					),
			),
		form.NewFieldCsrf("_csrf_token").WithData("my-token"),
		form.NewSubmit("submit").
			WithOptions(
				form.NewOption("attr", form.Attrs{
					"class": "btn btn-primary",
				}),
			),
	).
		End().
		WithOptions(
			form.NewOption("help", "form help"),
			form.NewOption("help_attr", form.Attrs{
				"class": "btn btn-primary",
			}),
			form.NewOption("attr", form.Attrs{
				"class": "row",
			}),
		).
		WithMethod(http.MethodPost).
		WithAction("/")
}
