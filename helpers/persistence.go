package helpers

// Persistent describes methods that can be applied with a fully persistent struct
type Persistent interface {
	// With mutates a single position to the given letter, returning a new object
	With(...interface{}) *Persistent
}
