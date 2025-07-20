package form

func NewFieldCsrf(name string) *Field {
	f := NewFieldHidden(name).
		WithFixedName()

	return f
}
