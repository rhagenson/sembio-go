package interfaces

// Sequence is an abstract type defining the basic functionality of any biological sequence (DNA, RNA, Protein, or some other series of characters in a given alphabet)
// string type is returned to implictly allow non-utf8, multi-character elements which would allow defining a k-mer Sequence.
type Sequence interface {
	// Length returns how many elements there are in the Sequence
	Length() int

	// Alphabet returns the valid elements in the Sequence
	Alphabet() []string

	// Position returns the n-th element
	Position(n int) string

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop int) []string
}
