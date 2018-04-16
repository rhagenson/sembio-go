package alphabet

// Ambiguouser implements the Ambiguous method, which tells the user if a struct contains an ambiguity
type Ambiguouser interface {
	// Ambiguous tells the user if the struct includes ambiguity
	Ambiguous() bool
}
