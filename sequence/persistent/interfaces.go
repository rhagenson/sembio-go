package persistent

// ErrorAccumulator keeps a record of past transformation errors on an struct
type ErrorAccumulator interface {
	// Errors returns any accumulated errors that result from chaining operations
	Errors() []error
}
