package sequence

import "bitbucket.org/rhagenson/bigr/alphabet"

// Interface is an abstract type defining the basic functionality of any biological sequence
// (DNA, RNA, Protein, or some other series of Letters in a given alphabet)
// A Interface is essentially an ordered list of elements where there is a Lenth(), each element has a Position(),
// it is possible to look at a specific Range() of elements
type Interface interface {
	// Length returns how many elements there are in the Sequence
	Length() uint

	// Position returns the n-th element
	Position(n uint) string

	// Range returns elements from start (inclusive) to stop (exclusive)
	Range(start, stop uint) string
}

// DNAstricter is any representation that uses the StrictDNAAlphabet, which
// only has four (4) letters that are possible.
type DNAstricter interface {
	Alphabet() alphabet.DNAStrict
}

// DNAiupacer is any representation that uses the IUPACDNAAlphabet, which
// has all sixteen (16) letters that are possible.
type DNAiupacer interface {
	Alphabet() alphabet.DNAIupac
}

// StrictDNASequence is the combination of implementing both
// StrictDNA and Sequence interfaces
type StrictDNASequence interface {
	DNAstricter
	Interface
}

// IupacDNASequence is the combination of implementing both
// IupacDNA and Sequence interfaces
type IupacDNASequence interface {
	DNAiupacer
	Interface
}

// RNAstricter is any representation that uses the StrictRNAAlphabet, which
// only has four (4) letters that are possible.
type RNAstricter interface {
	Alphabet() alphabet.RNAStrict
}

// RNAiupacer is any representation that uses the IUPACRNAAlphabet, which
// has all sixteen (16) letters that are possible.
type RNAiupacer interface {
	Alphabet() alphabet.RNAIupac
}

// StrictRNASequence is the combination of implementing both
// StrictRNA and Sequence interfaces
type StrictRNASequence interface {
	RNAstricter
	Interface
}

// IupacRNASequence is the combination of implementing both
// IupacRNA and Sequence interfaces
type IupacRNASequence interface {
	RNAiupacer
	Interface
}

// ProteinStricter is any representation that uses the ProteinStrict Alphabet
type ProteinStricter interface {
	Alphabet() alphabet.ProteinStrict
}

// ProteinGappeder is any representation that uses the ProteinGapped Alphabet
type ProteinGappeder interface {
	Alphabet() alphabet.ProteinGapped
}

// StrictProteinSequence is the combination of implementing both
// StrictProtein and Sequence interfaces
type StrictProteinSequence interface {
	ProteinStricter
	Interface
}

// GappedProteinSequence is the combination of implementing both
// GappedProtein and Sequence interfaces
type GappedProteinSequence interface {
	ProteinGappeder
	Interface
}
