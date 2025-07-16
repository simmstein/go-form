package validation

type Constraint interface {
	Validate(data any) []Error
}
