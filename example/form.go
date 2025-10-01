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
	Collection []CollectionItem
	Bytes      []byte
	Text       string
	Checkbox   bool
	Dates      ExampleDates
	Choices    ExampleChoices
	Inputs     ExampleOtherInputs
}

type CollectionItem struct {
	ValueA string
	ValueB string
}

type Theme struct {
	Value string `field:"lowerCamel"`
}

func CreateDataForm(action string) *form.Form {
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
			).
			WithConstraints(
				validation.NewRegex(`deblan`),
			),
		form.NewFieldCheckbox("Checkbox").
			WithOptions(
				form.NewOption("label", "Checkbox"),
			),
		form.NewSubForm("Inputs").
			WithOptions(
				form.NewOption("label", "Inputs"),
			).
			Add(
				form.NewFieldNumber("Number").
					WithOptions(
						form.NewOption("label", "Number"),
					).
					WithConstraints(
						validation.NewRange().WithRange(1, 20),
						validation.NewIsEven(),
					),
				form.NewFieldRange("Range").
					WithOptions(
						form.NewOption("label", "Range"),
					),
				form.NewFieldMail("Mail").
					WithOptions(
						form.NewOption("label", "Mail"),
					).
					WithConstraints(
						validation.Mail{},
					),
				form.NewFieldPassword("Password").
					WithOptions(
						form.NewOption("label", "Password"),
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
					),
				form.NewFieldDatetime("DateTime").
					WithOptions(
						form.NewOption("label", "Datetime"),
					),
				form.NewFieldDatetimeLocal("DateTimeLocal").
					WithOptions(
						form.NewOption("label", "DateTime local"),
					),
				form.NewFieldTime("Time").
					WithOptions(
						form.NewOption("label", "Time"),
					),
			),
		form.NewSubForm("Choices").
			WithOptions(form.NewOption("label", "Choices")).
			Add(
				form.NewFieldChoice("Select").
					WithOptions(
						form.NewOption("choices", itemsChoices),
						form.NewOption("label", "Select"),
					).
					WithConstraints(
						validation.NewNotBlank(),
					),
				form.NewFieldChoice("SelectExpanded").
					WithOptions(
						form.NewOption("choices", itemsChoices),
						form.NewOption("label", "Select (expanded)"),
						form.NewOption("expanded", true),
					),
				form.NewFieldChoice("MultipleSelect").
					WithSlice().
					WithOptions(
						form.NewOption("choices", itemsChoices),
						form.NewOption("label", "Multiple select"),
						form.NewOption("multiple", true),
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
					),
			),
		form.NewFieldCollection("Collection").
			WithOptions(
				form.NewOption("label", "Collection"),
				form.NewOption("form", form.NewForm(
					form.NewFieldText("ValueA").
						WithOptions(form.NewOption("label", "Value A")).
						WithConstraints(
							validation.NewNotBlank(),
						),
					form.NewFieldText("ValueB").
						WithOptions(form.NewOption("label", "Value B")),
				)),
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
			form.NewOption("help", "Form global help"),
		).
		WithMethod(http.MethodPost).
		WithAction(action)
}

func NewTheme(value string) *Theme {
	return &Theme{Value: value}
}

func CreateThemeSelectorForm() *form.Form {
	choices := form.NewChoices([]map[string]string{
		map[string]string{"value": "/", "label": "Html5"},
		map[string]string{"value": "/bootstrap", "label": "Bootstrap5"},
	})

	choices.LabelBuilder = func(key int, item any) string {
		return item.(map[string]string)["label"]
	}

	choices.ValueBuilder = func(key int, item any) string {
		return item.(map[string]string)["value"]
	}

	return form.NewForm(
		form.NewFieldChoice("value").
			WithOptions(
				form.NewOption("choices", choices),
				form.NewOption("label", "Select a theme"),
				form.NewOption("required", true),
				form.NewOption("attr", form.Attrs{
					"onchange": "document.location.href = this.value",
				}),
			),
	).
		End().
		WithName("").
		WithMethod(http.MethodGet)
}
