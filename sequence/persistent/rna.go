package persistent

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/helpers/complement"
	"bitbucket.org/rhagenson/bio/sequence"
)

var _ sequence.Reverser = new(Rna)
var _ sequence.RevComper = new(Rna)
var _ sequence.Complementer = new(Rna)
var _ Wither = new(Rna)

// Rna is a sequence witch validates against the Rna alphabet
// and knows how to reverse, complement, and revcomp itself
type Rna struct {
	*Struct
}

// NewRna generates a New sequence that validates against the Rna alphabet
func NewRna(s string) (*Rna, error) {
	n := NewStruct(
		s,
		AlphabetIs(alphabet.Rna),
	)
	return &Rna{n}, n.Validate()
}

// Reverse is the same Rna with the sequence reversed
func (x *Rna) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewRna(string(t))
}

// RevComp is the same Rna with the sequence reversed and complemented
func (x *Rna) RevComp() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.Rna(t[l-1-i]), complement.Rna(t[i])
	}
	return NewRna(string(t))
}

// Complement is the same Rna with the sequence complemented
func (x *Rna) Complement() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.Rna(byte(x.seq[i]))
	}
	return NewRna(string(t))
}
