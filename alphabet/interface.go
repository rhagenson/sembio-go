package alphabet

// Interface is the abstract representation of a standard N-letter alphabet
// where validity of a given Letter can be checked.
type Interface interface {
	// Contains checks that the given elements are in the Alphabet
	Contains([]byte) []bool
}
