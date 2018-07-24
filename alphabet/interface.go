package alphabet

// Interface is the abstract representation of a standard N-letter alphabet
// where validity of a given latter can be checked.
type Interface interface {
	// Contains checks that the given elements are in the Alphabet
	Contains(...string) []bool

	// Length is the number of letters in the Alphabet
	Length() int

	// String will return the equivalent string denoting valid letters
	String() string

	// Width is the byte width of letters
	// Zero (default int value) is the same as one to represent single-byte width
	Width() int
}
