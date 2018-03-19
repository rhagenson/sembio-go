package alphabet

// Alphabet is the abstract representation of a standard N-letter alphabet where validity of a given
// Letter can be checked.
type Alphabet interface {
	// Valid checks that a given Letter is in the Alphabet
	Valid(Letter) bool

	// Letters is all the valid Letter entries in the Alphabet
	Letters() []Letter

	// Length is the number of Letters in the Alphabet
	Length() int

	// Gapped is whether the Alphabet contains a gap character or not
	Gapped() bool

	// Ambiguous is whether the Alphabet contains an ambiguity character or not
	Ambiguous() bool
}
