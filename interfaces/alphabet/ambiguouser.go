package alphabet


// Ambiguouser implements the Ambiguous method, which tells the user if a struct contains an ambiguous Letter
type Ambiguouser interface {
	// Ambiguous tells the user if the struct includes an ambiguous character of some kind
	Ambiguous() bool
}
