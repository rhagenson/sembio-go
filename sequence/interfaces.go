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

// DnaStrictSequence is the combination of implementing both
// StrictDNA and Sequence interfaces
type DnaStrictSequence interface {
	Interface
	Alphabet() *alphabet.DNAStrict
}

// DnaIupacSequence is the combination of implementing both
// IupacDNA and Sequence interfaces
type DnaIupacSequence interface {
	Interface
	Alphabet() *alphabet.DNAIupac
}

// ProteinStrictSequence is the combination of implementing both
// StrictProtein and Sequence interfaces
type ProteinStrictSequence interface {
	Interface
	Alphabet() *alphabet.ProteinStrict
}

// ProteinGappedSequence is the combination of implementing both
// GappedProtein and Sequence interfaces
type ProteinGappedSequence interface {
	Interface
	Alphabet() *alphabet.ProteinGapped
}

// RnaStrictSequence is the combination of implementing both
// StrictRNA and Sequence interfaces
type RnaStrictSequence interface {
	Interface
	Alphabet() *alphabet.RNAStrict
}

// RnaIupacSequence is the combination of implementing both
// IupacRNA and Sequence interfaces
type RnaIupacSequence interface {
	Interface
	Alphabet() *alphabet.RNAIupac
}
