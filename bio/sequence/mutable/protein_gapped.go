package mutable

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/alphabet/hashmap"
	"github.com/rhagenson/bio-go/bio/sequence"
)

var _ sequence.Reverser = new(ProteinGapped)
var _ sequence.Alphabeter = new(ProteinGapped)
var _ sequence.LetterCounter = new(ProteinGapped)
var _ sequence.Validator = new(ProteinGapped)
var _ Wither = new(ProteinGapped)

// ProteinGapped is a sequence which validates against the ProteinGapped alphabet
// and knows how to reverse itself
type ProteinGapped struct {
	*Struct
}

// NewProteinGapped generates a New sequence that validates against the ProteinGapped alphabet
func NewProteinGapped(s string) (*ProteinGapped, error) {
	n := New(
		s,
		sequence.AlphabetIs(hashmap.NewProteinGapped()),
	)
	return &ProteinGapped{n}, n.Validate()
}

// Reverse is the same ProteinGapped with the sequence reversed
func (x *ProteinGapped) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	x.seq = string(t)
	return x, x.Validate()
}

// Alphabet reveals the underlying alphabet in use
func (x *ProteinGapped) Alphabet() alphabet.Interface {
	return hashmap.NewProteinGapped()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *ProteinGapped) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
