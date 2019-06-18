package mutable

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/sequence"
	"github.com/rhagenson/bio-go/bio/utils/complement"
)

var _ sequence.Reverser = new(RnaIupac)
var _ sequence.RevComper = new(RnaIupac)
var _ sequence.Complementer = new(RnaIupac)
var _ sequence.Alphabeter = new(RnaIupac)
var _ sequence.LetterCounter = new(RnaIupac)
var _ Wither = new(RnaIupac)
var _ Validator = new(RnaIupac)

// RnaIupac is a sequence witch validates against the RnaIupac alphabet
// and knows how to reverse, complement, and revcomp itself
type RnaIupac struct {
	*Struct
}

// NewRnaIupac generates a New sequence that validates against the RnaIupac alphabet
func NewRnaIupac(s string) (*RnaIupac, error) {
	n := New(
		s,
		AlphabetIs(alphabet.RnaIupac),
	)
	return &RnaIupac{n}, n.Validate()
}

// Reverse is the same RnaIupac with the sequence reversed
func (x *RnaIupac) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	x.seq = string(t)
	return x, x.Validate()
}

// RevComp is the same RnaIupac with the sequence reversed and complemented
func (x *RnaIupac) RevComp() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.RnaIupac(t[l-1-i]), complement.RnaIupac(t[i])
	}
	x.seq = string(t)
	return x, x.Validate()
}

// Complement is the same RnaIupac with the sequence complemented
func (x *RnaIupac) Complement() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.RnaIupac(t[i])
	}
	x.seq = string(t)
	return x, x.Validate()
}

// Alphabet reveals the underlying alphabet in use
func (x *RnaIupac) Alphabet() alphabet.Interface {
	return alphabet.RnaIupac
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *RnaIupac) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
