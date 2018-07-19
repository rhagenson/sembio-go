package sequence

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
	Reverse() (*Sequence, error)
}

// Complementer can complement the sequence
type Complementer interface {
	Complement() (*Sequence, error)
}

// RevComper can reverse-complement the sequence
type RevComper interface {
	RevComp() (*Sequence, error)
}

// Translater can translate the sequence
type Translater interface {
	Translate() (*Sequence, error)
}

// Transcriber can transcribe the sequence
type Transcriber interface {
	Transcribe() (*Sequence, error)
}
