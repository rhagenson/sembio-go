package sequence

import "bitbucket.org/rhagenson/bigr/alphabet"

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

// Alphabeter provides a way to get at the sequence valid alphabet
type Alphabeter interface {
	Alphabet() alphabet.Interface
}

// Complementer provides a way to get at the sequence complement
type Complementer interface {
	Complement() Interface
}

// Reverser provides a way to get at the reverse sequence
type Reverser interface {
	Reverse() Interface
}

// RevComper provides a way to get at the reverse complement directly
type RevComper interface {
	RevComp() Interface
}
