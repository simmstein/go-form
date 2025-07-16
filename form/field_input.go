package form

func NewFieldText(name string) *Field {
	f := NewField(name, "input").
		WithOptions(Option{Name: "type", Value: "text"})

	return f
}

func NewFieldNumber(name string) *Field {
	f := NewField(name, "input").
		WithOptions(Option{Name: "type", Value: "number"})

	return f
}

func NewSubmit(name string) *Field {
	f := NewField(name, "input").
		WithOptions(
			Option{Name: "type", Value: "submit"},
		)

	f.Data = "Submit"

	return f
}
