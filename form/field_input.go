package form

func NewFieldText(name string) *Field {
	f := NewField(name, "input").
		WithOptions(Option{Name: "type", Value: "text"})

	return f
}
