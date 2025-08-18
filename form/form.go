package form

// @license GNU AGPL version 3 or any later version
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import (
	"encoding/json"
	"io/ioutil"
	"maps"
	"net/http"
	"net/url"
	"slices"

	"github.com/mitchellh/mapstructure"
	"gitnet.fr/deblan/go-form/util"
	"gitnet.fr/deblan/go-form/validation"
)

// Field represents a form
type Form struct {
	Fields       []*Field
	GlobalFields []*Field
	Errors       []validation.Error
	Method       string
	JsonRequest  bool
	Action       string
	Name         string
	Options      []*Option
	RequestData  *url.Values
}

// Generates a new form with default properties
func NewForm(fields ...*Field) *Form {
	f := new(Form)
	f.Method = "POST"
	f.Name = "form"
	f.Add(fields...)
	f.WithOptions(
		NewOption("attr", Attrs{}),
		NewOption("help_attr", Attrs{}),
	)

	return f
}

// Checks if the form contains an option using its name
func (f *Form) HasOption(name string) bool {
	for _, option := range f.Options {
		if option.Name == name {
			return true
		}
	}

	return false
}

// Returns an option using its name
func (f *Form) GetOption(name string) *Option {
	for _, option := range f.Options {
		if option.Name == name {
			return option
		}
	}

	return nil
}

// Resets the form errors
func (f *Form) ResetErrors() *Form {
	f.Errors = []validation.Error{}

	return f
}

// Appends children
func (f *Form) Add(fields ...*Field) {
	for _, field := range fields {
		field.Form = f
		f.Fields = append(f.Fields, field)
	}
}

// Configures its children deeply
// This function must be called after adding all fields
func (f *Form) End() *Form {
	f.GlobalFields = []*Field{}

	for _, c := range f.Fields {
		f.AddGlobalField(c)
	}

	return f
}

// Configures its children deeply
func (f *Form) AddGlobalField(field *Field) {
	f.GlobalFields = append(f.GlobalFields, field)

	for _, c := range field.Children {
		f.AddGlobalField(c)
	}
}

// Checks if the form contains a child using its name
func (f *Form) HasField(name string) bool {
	for _, field := range f.Fields {
		if name == field.Name {
			return true
		}
	}

	return false
}

// Returns a child using its name
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

// Sets the method of the format (http.MethodPost, http.MethodGet, ...)
func (f *Form) WithMethod(v string) *Form {
	f.Method = v

	return f
}

// Sets the name of the form (used to compute name of fields)
func (f *Form) WithName(v string) *Form {
	f.Name = v

	return f
}

// Sets the action of the form (eg: "/")
func (f *Form) WithAction(v string) *Form {
	f.Action = v

	return f
}

// Appends options to the form
func (f *Form) WithOptions(options ...*Option) *Form {
	for _, option := range options {
		if f.HasOption(option.Name) {
			f.GetOption(option.Name).Value = option.Value
		} else {
			f.Options = append(f.Options, option)
		}
	}

	return f
}

// Checks the a form is valid
func (f *Form) IsValid() bool {
	isValid := true
	f.ResetErrors()

	for _, field := range f.Fields {
		fieldIsValid := field.Validate(field)
		isValid = isValid && fieldIsValid
	}

	return isValid
}

// Copies datas from a struct to the form
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

func (f *Form) WithJsonRequest() *Form {
	f.JsonRequest = true

	return f
}

// Copies datas from the form to a struct
func (f *Form) Bind(data any) error {
	toBind := make(map[string]any)

	for _, field := range f.Fields {
		field.Bind(toBind, nil)
	}

	return mapstructure.Decode(toBind, data)
}

// Processes a request
func (f *Form) HandleRequest(req *http.Request) {
	var data url.Values

	if f.JsonRequest {
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			return
		}

		mapping := make(map[string]any)
		err = json.Unmarshal(body, &mapping)

		if err != nil {
			return
		}

		data = url.Values{}
		util.MapToUrlValues(&data, f.Name, mapping)
	} else {
		switch f.Method {
		case "GET":
			data = req.URL.Query()
		default:
			req.ParseForm()
			data = req.Form
		}
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

// Checks if the form is submitted
func (f *Form) IsSubmitted() bool {
	return f.RequestData != nil
}

func (f *Form) ErrorsTree() map[string]any {
	errors := make(map[string]any)

	for _, field := range f.Fields {
		field.ErrorsTree(errors, nil)
	}

	return map[string]any{
		"errors":   f.Errors,
		"children": slices.Collect(maps.Values(errors)),
	}
}
