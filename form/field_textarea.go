package form

// Generates a textarea
func NewFieldTextarea(name string) *Field {
	return NewField(name, "textarea")
}
