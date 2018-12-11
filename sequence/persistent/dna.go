package persistent

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/helpers/complement"
	"bitbucket.org/rhagenson/bio/sequence"
)

var _ sequence.Reverser = new(Dna)
var _ sequence.RevComper = new(Dna)
var _ sequence.Complementer = new(Dna)
var _ Wither = new(Dna)

// Dna is a sequence witch validates against the Dna alphabet
// and knows how to reverse, complement, and revcomp itself
type Dna struct {
	*Struct
}

// NewDna generates a New sequence that validates against the Dna alphabet
func NewDna(s string) (*Dna, error) {
	n := NewStruct(
		s,
		AlphabetIs(alphabet.Dna),
	)
	return &Dna{n}, n.Validate()
}

// Reverse is the same Dna with the sequence reversed
func (x *Dna) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewDna(string(t))
}

// RevComp is the same Dna with the sequence reversed and complemented
func (x *Dna) RevComp() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.Dna(t[l-1-i]), complement.Dna(t[i])
	}
	return NewDna(string(t))
}

// Complement is the same DnaIupac with the sequence complemented
func (x *Dna) Complement() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.Dna(byte(x.seq[i]))
	}
	return NewDna(string(t))
}
