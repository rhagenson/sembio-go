package alphabet

// Alphabet is the abstract representation of a standard N-letter alphabet where validity of a given
// Letter can be checked.
// As well, the Alphabet must be able return an array of all its valid Letter elements, be gapped or not, and be
// ambiguous or not.
type Alphabet interface {
	// Contains checks that given Letter elements are in the Alphabet
	Contains(...Letter) []bool

	// Letters is all valid Letter elements in the Alphabet (be sure to return a copy not a reference slice)
	Letters() []Letter

	// An Alphabet must either implement gaps or not
	Gapper

	// An Alphabet must either implement ambiguous letters or not
	Ambiguouser
}
