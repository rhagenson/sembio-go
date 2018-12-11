package sequence

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/helpers/complement"
)

var _ Reverser = new(Rna)
var _ RevComper = new(Rna)
var _ Complementer = new(Rna)

// Rna is a sequence witch validates aginst the Rna alphabet
// and knows how to reverse, complement, and revcomp itself
type Rna struct {
	*backer
}

// NewRna generates a New sequence that validates against the Rna alphabet
func NewRna(s string) (*Rna, error) {
	n := newBacker(
		s,
		AlphabetIs(alphabet.Rna),
	)
	return &Rna{n}, n.Validate()
}

// Reverse is the same Rna with the sequence reversed
func (x *Rna) Reverse() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewRna(string(t))
}

// RevComp is the same Rna with the sequence reversed and complemented
func (x *Rna) RevComp() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.Rna(t[l-1-i]), complement.Rna(t[i])
	}
	return NewRna(string(t))
}

// Complement is the same Rna with the sequence complemented
func (x *Rna) Complement() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.Rna(byte(x.seq[i]))
	}
	return NewRna(string(t))
}
