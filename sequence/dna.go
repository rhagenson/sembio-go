package sequence

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/helpers/complement"
)

var _ Reverser = new(Dna)
var _ RevComper = new(Dna)
var _ Complementer = new(Dna)

// Dna is a sequence witch validates aginst the Dna alphabet
// and knows how to reverse, complement, and revcomp itself
type Dna struct {
	*backer
}

// NewDna generates a New sequence that validates against the Dna alphabet
func NewDna(s string) (*Dna, error) {
	n := newBacker(
		s,
		AlphabetIs(alphabet.Dna),
	)
	return &Dna{n}, n.Validate()
}

// Reverse is the same Dna with the sequence reversed
func (x *Dna) Reverse() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewDna(string(t))
}

// RevComp is the same Dna with the sequence reversed and complemented
func (x *Dna) RevComp() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.Dna(t[l-1-i]), complement.Dna(t[i])
	}
	return NewDna(string(t))
}

// Complement is the same DnaIupac with the sequence complemented
func (x *Dna) Complement() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.Dna(byte(x.seq[i]))
	}
	return NewDna(string(t))
}
