package form

import (
	"net/http"

	"gitnet.fr/deblan/go-form/util"
	"gitnet.fr/deblan/go-form/validation"
)

type Form struct {
	Fields  []*Field
	Errors  []validation.Error
	Method  string
	Action  string
	Name    string
	Options []Option
}

func NewForm(fields ...*Field) *Form {
	f := new(Form)
	f.Method = "POST"
	f.Add(fields...)

	return f
}

func (f *Form) Add(fields ...*Field) {
	for _, field := range fields {
		f.Fields = append(f.Fields, field)
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

func (f *Form) WithOptions(options ...Option) *Form {
	for _, option := range options {
		f.Options = append(f.Options, option)
	}

	return f
}

func (f *Form) IsValid() bool {
	isValid := true

	for _, field := range f.Fields {
		fieldIsValid := field.Validate(field)
		isValid = isValid && fieldIsValid
	}

	return isValid
}

func (f *Form) Bind(data any) error {
	props, err := util.InspectStruct(data)

	if err != nil {
		return err
	}

	for key, value := range props {
		if f.HasField(key) {
			err = f.GetField(key).Bind(value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (f *Form) HandleRequest(req http.Request) {
	if f.Method == "POST" {
		// data := req.PostForm
	}
}
