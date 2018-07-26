package alphabet

import "fmt"

// Interface is the abstract representation of a finite-letter alphabet
// where validity of a given letter can be checked.
type Interface interface {
	// Must be stringable
	fmt.Stringer

	// Contains checks that the given elements are in the Alphabet
	Contains(...string) []bool

	// Length is the number of letters in the Alphabet
	Length() int

	// Width is the byte width of letters
	// Zero (default int value) is the same as one to represent single-byte width
	Width() uint
}
