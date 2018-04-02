package alphabet

// Alphabet is the abstract representation of a standard N-letter alphabet where validity of a given
// Letter can be checked.
type Alphabet interface {
	// Contains checks that given Letter elements are in the Alphabet
	Contains(...Letter) []bool
}
