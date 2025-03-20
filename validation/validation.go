package validation

type ValidationError struct {
	Err          error
	Field        string
	ErrorMessage string
}

func (e *ValidationError) Error() string { return e.Err.Error() }
func (e *ValidationError) Unwrap() error { return e.Err }
