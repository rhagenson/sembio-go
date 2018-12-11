package sequence

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/helpers/complement"
)

var _ Reverser = new(RnaIupac)
var _ RevComper = new(RnaIupac)
var _ Complementer = new(RnaIupac)

// RnaIupac is a sequence witch validates aginst the RnaIupac alphabet
// and knows how to reverse, complement, and revcomp itself
type RnaIupac struct {
	*Struct
}

// NewRnaIupac generates a New sequence that validates against the RnaIupac alphabet
func NewRnaIupac(s string) (*RnaIupac, error) {
	n := NewStruct(
		s,
		AlphabetIs(alphabet.RnaIupac),
	)
	return &RnaIupac{n}, n.Validate()
}

// Reverse is the same RnaIupac with the sequence reversed
func (x *RnaIupac) Reverse() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewRnaIupac(string(t))
}

// RevComp is the same RnaIupac with the sequence reversed and complemented
func (x *RnaIupac) RevComp() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.RnaIupac(t[l-1-i]), complement.RnaIupac(t[i])
	}
	return NewRnaIupac(string(t))
}

// Complement is the same RnaIupac with the sequence complemented
func (x *RnaIupac) Complement() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.RnaIupac(t[i])
	}
	return NewRnaIupac(string(t))
}
