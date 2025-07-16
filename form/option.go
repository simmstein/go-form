package form

type Option struct {
	Name  string
	Value any
}

func NewOption(name string, value any) *Option {
	return &Option{
		Name:  name,
		Value: value,
	}
}
