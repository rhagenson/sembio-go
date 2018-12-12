package sequence

import (
	"bitbucket.org/rhagenson/bio/data/codon"
)

// Interface is the basic functionality of any biological sequence
// (DNA, RNA, Protein, or other)
type Interface interface {
	// Length is the number of elements in the Interface
	Length() uint

	// Position is the n-th element
	Position(n uint) (string, error)

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop uint) (string, error)
}

// Reverser can reverse the sequence
type Reverser interface {
	Reverse() (Interface, error)
}

// Complementer can complement the sequence
type Complementer interface {
	Complement() (Interface, error)
}

// RevComper can reverse-complement the sequence
type RevComper interface {
	RevComp() (Interface, error)
}

// Translater can translate the sequence
type Translater interface {
	Translate(codon.Translater) (Interface, error)
}

// Transcriber can transcribe the sequence
type Transcriber interface {
	Transcribe() (Interface, error)
}
