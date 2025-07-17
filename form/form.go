package form

import (
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
	"gitnet.fr/deblan/go-form/util"
	"gitnet.fr/deblan/go-form/validation"
)

type Form struct {
	Fields       []*Field
	GlobalFields []*Field
	Errors       []validation.Error
	Method       string
	Action       string
	Name         string
	Options      []*Option
	RequestData  *url.Values
}

func NewForm(fields ...*Field) *Form {
	f := new(Form)
	f.Method = "POST"
	f.Name = "form"
	f.Add(fields...)

	return f
}

func (f *Form) HasOption(name string) bool {
	for _, option := range f.Options {
		if option.Name == name {
			return true
		}
	}

	return false
}

func (f *Form) GetOption(name string) *Option {
	for _, option := range f.Options {
		if option.Name == name {
			return option
		}
	}

	return nil
}

func (f *Form) ResetErrors() *Form {
	f.Errors = []validation.Error{}

	return f
}

func (f *Form) Add(fields ...*Field) {
	for _, field := range fields {
		field.Form = f
		f.Fields = append(f.Fields, field)
	}
}

func (f *Form) End() *Form {
	for _, c := range f.Fields {
		f.AddGlobalField(c)
	}

	return f
}

func (f *Form) AddGlobalField(field *Field) {
	f.GlobalFields = append(f.GlobalFields, field)

	for _, c := range field.Children {
		f.AddGlobalField(c)
	}
}

func (f *Form) HasField(name string) bool {
	for _, field := range f.Fields {
		if name == field.Name {
			return true
		}
	}

	return false
}

func (f *Form) GetField(name string) *Field {
	var result *Field

	for _, field := range f.Fields {
		if name == field.Name {
			result = field
			break
		}
	}

	return result
}

func (f *Form) WithMethod(v string) *Form {
	f.Method = v

	return f
}

func (f *Form) WithName(v string) *Form {
	f.Name = v

	return f
}

func (f *Form) WithAction(v string) *Form {
	f.Action = v

	return f
}

func (f *Form) WithOptions(options ...*Option) *Form {
	for _, option := range options {
		f.Options = append(f.Options, option)
	}

	return f
}

func (f *Form) IsValid() bool {
	isValid := true
	f.ResetErrors()

	for _, field := range f.Fields {
		fieldIsValid := field.Validate(field)
		isValid = isValid && fieldIsValid
	}

	return isValid
}

func (f *Form) Mount(data any) error {
	props, err := util.InspectStruct(data)

	if err != nil {
		return err
	}

	for key, value := range props {
		if f.HasField(key) {
			err = f.GetField(key).Mount(value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (f *Form) Bind(data any) error {
	toBind := make(map[string]any)

	for _, field := range f.Fields {
		field.Bind(toBind, nil)
	}

	return mapstructure.Decode(toBind, data)
}

func (f *Form) HandleRequest(req *http.Request) {
	var data url.Values

	if f.Method != "GET" {
		req.ParseForm()
		data = req.Form
	} else {
		data = req.URL.Query()
	}

	isSubmitted := false

	for _, c := range f.GlobalFields {
		if data.Has(c.GetName()) {
			isSubmitted = true

			if c.IsSlice {
				c.Mount(data[c.GetName()])
			} else {
				c.Mount(data.Get(c.GetName()))
			}
		}
	}

	if isSubmitted {
		f.RequestData = &data
	}
}

func (f *Form) IsSubmitted() bool {
	return f.RequestData != nil
}
