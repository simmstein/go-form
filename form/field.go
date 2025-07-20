package form

import (
	"fmt"
	"strings"

	"gitnet.fr/deblan/go-form/util"
	"gitnet.fr/deblan/go-form/validation"
)

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

type Field struct {
	Name        string
	Widget      string
	Data        any
	Options     []*Option
	Children    []*Field
	Constraints []validation.Constraint
	Errors      []validation.Error
	PrepareView func() map[string]any
	BeforeMount func(data any) (any, error)
	BeforeBind  func(data any) (any, error)
	Validate    func(f *Field) bool
	IsSlice     bool
	IsFixedName bool
	Form        *Form
	Parent      *Field
}

func NewField(name, widget string) *Field {
	f := &Field{
		Name:        name,
		IsFixedName: false,
		Widget:      widget,
		Data:        nil,
	}

	f.PrepareView = func() map[string]any {
		m := make(map[string]any)

		return m
	}

	f.BeforeMount = func(data any) (any, error) {
		return data, nil
	}

	f.BeforeBind = func(data any) (any, error) {
		return data, nil
	}

	f.Validate = FieldValidation

	return f
}

func (f *Field) HasOption(name string) bool {
	for _, option := range f.Options {
		if option.Name == name {
			return true
		}
	}

	return false
}

func (f *Field) GetOption(name string) *Option {
	for _, option := range f.Options {
		if option.Name == name {
			return option
		}
	}

	return nil
}

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

func (f *Field) WithData(data any) *Field {
	f.Data = data

	return f
}

func (f *Field) ResetErrors() *Field {
	f.Errors = []validation.Error{}

	return f
}

func (f *Field) WithSlice() *Field {
	f.IsSlice = true

	return f
}

func (f *Field) WithFixedName() *Field {
	f.IsFixedName = true

	return f
}

func (f *Field) WithConstraints(constraints ...validation.Constraint) *Field {
	for _, constraint := range constraints {
		f.Constraints = append(f.Constraints, constraint)
	}

	return f
}

func (f *Field) WithBeforeMount(callback func(data any) (any, error)) *Field {
	f.BeforeMount = callback

	return f
}

func (f *Field) WithBeforeBind(callback func(data any) (any, error)) *Field {
	f.BeforeBind = callback

	return f
}

func (f *Field) Add(children ...*Field) *Field {
	for _, child := range children {
		child.Parent = f
		f.Children = append(f.Children, child)
	}

	return f
}

func (f *Field) HasChild(name string) bool {
	for _, child := range f.Children {
		if name == child.Name {
			return true
		}
	}

	return false
}

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

func (f *Field) GetId() string {
	name := f.GetName()
	name = strings.ReplaceAll(name, "[", "-")
	name = strings.ReplaceAll(name, "]", "")
	name = strings.ToLower(name)

	return name
}

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
