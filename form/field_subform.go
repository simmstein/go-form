package form

func NewFieldSubForm(name string) *Field {
	f := NewField(name, "sub_form")

	return f
}

func NewSubForm(name string) *Field {
	return NewFieldSubForm(name)
}
