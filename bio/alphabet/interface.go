package alphabet

import "fmt"

// Interface is the abstract representation of a finite-letter alphabet
// where validity of a given letter can be checked.
type Interface interface {
	// Must be stringable
	fmt.Stringer

	// Contains checks that the given elements are in the Alphabet
	Contains(...byte) []bool

	// Length is the number of letters in the Alphabet
	Length() int
}

// Complementer is any alphabet that has complementing characters
type Complementer interface {
	Complement(byte) byte
}
