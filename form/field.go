package form

import (
	"gitnet.fr/deblan/go-form/util"
	"gitnet.fr/deblan/go-form/validation"
)

func FieldValidation(f *Field) bool {
	if len(f.Children) > 0 {
		isValid := true

		for _, c := range f.Children {
			isChildValid, errs := validation.Validate(c.Data, c.Constraints)

			if len(errs) > 0 {
				c.Errors = errs
			}

			isValid = isValid && isChildValid
		}

		return isValid
	} else {
		isValid, errs := validation.Validate(f.Data, f.Constraints)
		f.Errors = []validation.Error{}

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
	BeforeBind  func(data any) (any, error)
	Validate    func(f *Field) bool
}

func NewField(name, widget string) *Field {
	f := &Field{
		Name:   name,
		Widget: widget,
		Data:   nil,
	}

	f.PrepareView = func() map[string]any {
		m := make(map[string]any)

		return m
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

func (f *Field) WithOptions(options ...Option) *Field {
	for _, option := range options {
		if f.HasOption(option.Name) {
			f.GetOption(option.Name).Value = option.Value
		} else {
			f.Options = append(f.Options, &option)
		}
	}

	return f
}

func (f *Field) WithConstraints(constraints ...validation.Constraint) *Field {
	for _, constraint := range constraints {
		f.Constraints = append(f.Constraints, constraint)
	}

	return f
}

func (f *Field) Add(children ...*Field) *Field {
	for _, child := range children {
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

func (f *Field) Bind(data any) error {
	if len(f.Children) == 0 {
		f.Data = data

		return nil
	}

	data, err := f.BeforeBind(data)

	if err != nil {
		return err
	}

	props, err := util.InspectStruct(data)

	if err != nil {
		return err
	}

	for key, value := range props {
		if f.HasChild(key) {
			err = f.GetChild(key).Bind(value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
