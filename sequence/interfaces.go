package sequence

// Interface is the basic functionality of any biological sequence
// (DNA, RNA, Protein, or other)
type Interface interface {
	// Length is the number of elements in the Interface
	Length() uint

	// Position is the n-th element
	Position(n uint) string

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop uint) string
}
