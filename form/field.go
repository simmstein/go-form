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
	"fmt"
	"strings"

	"gitnet.fr/deblan/go-form/util"
	"gitnet.fr/deblan/go-form/validation"
)

// Generic function for field.Validation
func FieldValidation(f *Field) bool {
	if len(f.Children) > 0 {
		isValid := true

		for _, c := range f.Children {
			c.ResetErrors()
			isChildValid, errs := validation.Validate(c.Data, c.Constraints)

			if len(errs) > 0 {
				c.Errors = errs
			}

			isValid = isValid && isChildValid
		}

		return isValid
	} else {
		f.ResetErrors()
		isValid, errs := validation.Validate(f.Data, f.Constraints)

		if len(errs) > 0 {
			f.Errors = errs
		}

		return isValid
	}
}

// Field represents a field in a form
type Field struct {
	Name        string
	Widget      string
	Data        any
	Options     []*Option
	Children    []*Field
	Constraints []validation.Constraint
	Errors      []validation.Error
	BeforeMount func(data any) (any, error)
	BeforeBind  func(data any) (any, error)
	Validate    func(f *Field) bool
	IsSlice     bool
	IsFixedName bool
	Form        *Form
	Parent      *Field
}

// Generates a new field with default properties
// It should not be used directly but inside function like in form.NewFieldText
func NewField(name, widget string) *Field {
	f := &Field{
		Name:        name,
		IsFixedName: false,
		Widget:      widget,
		Data:        nil,
	}

	f.BeforeMount = func(data any) (any, error) {
		return data, nil
	}

	f.BeforeBind = func(data any) (any, error) {
		return data, nil
	}

	f.WithOptions(
		NewOption("attr", Attrs{}),
		NewOption("row_attr", Attrs{}),
		NewOption("label_attr", Attrs{}),
		NewOption("help_attr", Attrs{}),
	)

	f.Validate = FieldValidation

	return f
}

// Checks if the field contains an option using its name
func (f *Field) HasOption(name string) bool {
	for _, option := range f.Options {
		if option.Name == name {
			return true
		}
	}

	return false
}

// Returns an option using its name
func (f *Field) GetOption(name string) *Option {
	for _, option := range f.Options {
		if option.Name == name {
			return option
		}
	}

	return nil
}

// Appends options to the field
func (f *Field) WithOptions(options ...*Option) *Field {
	for _, option := range options {
		if f.HasOption(option.Name) {
			f.GetOption(option.Name).Value = option.Value
		} else {
			f.Options = append(f.Options, option)
		}
	}

	return f
}

// Sets data the field
func (f *Field) WithData(data any) *Field {
	f.Data = data

	return f
}

// Resets the field errors
func (f *Field) ResetErrors() *Field {
	f.Errors = []validation.Error{}

	return f
}

// Sets that the field represents a data slice
func (f *Field) WithSlice() *Field {
	f.IsSlice = true

	return f
}

// Sets that the name of the field is not computed
func (f *Field) WithFixedName() *Field {
	f.IsFixedName = true

	return f
}

// Appends constraints
func (f *Field) WithConstraints(constraints ...validation.Constraint) *Field {
	for _, constraint := range constraints {
		f.Constraints = append(f.Constraints, constraint)
	}

	return f
}

// Sets a transformer applied to the structure data before displaying it in a field
func (f *Field) WithBeforeMount(callback func(data any) (any, error)) *Field {
	f.BeforeMount = callback

	return f
}

// Sets a transformer applied to the data of a field before defining it in a structure
func (f *Field) WithBeforeBind(callback func(data any) (any, error)) *Field {
	f.BeforeBind = callback

	return f
}

// Appends children
func (f *Field) Add(children ...*Field) *Field {
	for _, child := range children {
		child.Parent = f
		f.Children = append(f.Children, child)
	}

	return f
}

// Checks if the field contains a child using its name
func (f *Field) HasChild(name string) bool {
	for _, child := range f.Children {
		if name == child.Name {
			return true
		}
	}

	return false
}

// Returns a child using its name
func (f *Field) GetChild(name string) *Field {
	var result *Field

	for _, child := range f.Children {
		if name == child.Name {
			result = child

			break
		}
	}

	return result
}

// Computes the name of the field
func (f *Field) GetName() string {
	var name string

	if f.IsFixedName {
		return f.Name
	}

	if f.Form != nil && f.Form.Name != "" {
		name = fmt.Sprintf("%s[%s]", f.Form.Name, f.Name)
	} else if f.Parent != nil {
		name = fmt.Sprintf("%s[%s]", f.Parent.GetName(), f.Name)
	} else {
		name = f.Name
	}

	return name
}

// Computes the id of the field
func (f *Field) GetId() string {
	name := f.GetName()
	name = strings.ReplaceAll(name, "[", "-")
	name = strings.ReplaceAll(name, "]", "")
	name = strings.ToLower(name)

	return name
}

// Populates the field with data
func (f *Field) Mount(data any) error {
	data, err := f.BeforeMount(data)

	if err != nil {
		return err
	}

	if len(f.Children) == 0 {
		f.Data = data

		return nil
	}

	props, err := util.InspectStruct(data)

	if err != nil {
		return err
	}

	for key, value := range props {
		if f.HasChild(key) {
			err = f.GetChild(key).Mount(value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Bind the data into the given map
func (f *Field) Bind(data map[string]any, key *string) error {
	if len(f.Children) == 0 {
		v, err := f.BeforeBind(f.Data)

		if err != nil {
			return err
		}

		if key != nil {
			data[*key] = v
		} else {
			data[f.Name] = v
		}

		return nil
	}

	data[f.Name] = make(map[string]any)

	for _, child := range f.Children {
		child.Bind(data[f.Name].(map[string]any), key)
	}

	return nil
}

func (f *Field) ErrorsTree(tree map[string]any, key *string) {
	var index string

	if key != nil {
		index = *key
	} else {
		index = f.Name
	}

	if len(f.Children) == 0 {
		if len(f.Errors) > 0 {
			tree[index] = map[string]any{
				"meta": map[string]any{
					"id":       f.GetId(),
					"name":     f.Name,
					"formName": f.GetName(),
				},
				"errors": f.Errors,
			}
		}
	} else {
		errors := make(map[string]any)

		for _, child := range f.Children {
			if len(child.Errors) > 0 {
				child.ErrorsTree(errors, &child.Name)
			}
		}

		if len(errors) > 0 {
			tree[index] = map[string]any{
				"meta": map[string]any{
					"id":       f.GetId(),
					"name":     f.Name,
					"formName": f.GetName(),
				},
				"errors":   []validation.Error{},
				"children": errors,
			}
		}
	}
}
