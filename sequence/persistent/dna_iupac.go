package persistent

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/helpers/complement"
	"bitbucket.org/rhagenson/bio/sequence"
)

var _ sequence.Reverser = new(DnaIupac)
var _ sequence.RevComper = new(DnaIupac)
var _ sequence.Complementer = new(DnaIupac)
var _ Wither = new(DnaIupac)

// DnaIupac is a sequence witch validates aginst the DnaIupac alphabet
// and knows how to reverse, complement, and revcomp itself
type DnaIupac struct {
	*Struct
}

// NewDnaIupac generates a New sequence that validates against the DnaIupac alphabet
func NewDnaIupac(s string) (*DnaIupac, error) {
	n := NewStruct(
		s,
		AlphabetIs(alphabet.DnaIupac),
	)
	return &DnaIupac{n}, n.Validate()
}

// Reverse is the same DnaIupac with the sequence reversed
func (x *DnaIupac) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewDnaIupac(string(t))
}

// RevComp is the same DnaIupac with the sequence reversed and complemented
func (x *DnaIupac) RevComp() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.DnaIupac(t[l-1-i]), complement.DnaIupac(t[i])
	}
	return NewDnaIupac(string(t))
}

// Complement is the same DnaIupac with the sequence complemented
func (x *DnaIupac) Complement() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.DnaIupac(t[i])
	}
	return NewDnaIupac(string(t))
}