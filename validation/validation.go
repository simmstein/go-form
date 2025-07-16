package validation

func Validate(data any, constraints []Constraint) (bool, []Error) {
	errs := []Error{}

	for _, constraint := range constraints {
		for _, e := range constraint.Validate(data) {
			errs = append(errs, e)
		}
	}

	return len(errs) == 0, errs
}
