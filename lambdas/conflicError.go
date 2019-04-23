package lambdas

// ConflicError conflicts exceptions
type ConflictError struct {
}

func (e ConflictError) Error() string {
	return "ConflictError"
}
