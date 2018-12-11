package sequence

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/helpers/complement"
)

var _ Reverser = new(DnaIupac)
var _ RevComper = new(DnaIupac)
var _ Complementer = new(DnaIupac)

// DnaIupac is a sequence witch validates aginst the DnaIupac alphabet
// and knows how to reverse, complement, and revcomp itself
type DnaIupac struct {
	*backer
}

// NewDnaIupac generates a New sequence that validates against the DnaIupac alphabet
func NewDnaIupac(s string) (*DnaIupac, error) {
	n := newBacker(
		s,
		AlphabetIs(alphabet.DnaIupac),
	)
	return &DnaIupac{n}, n.Validate()
}

// Reverse is the same DnaIupac with the sequence reversed
func (x *DnaIupac) Reverse() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewDnaIupac(string(t))
}

// RevComp is the same DnaIupac with the sequence reversed and complemented
func (x *DnaIupac) RevComp() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.DnaIupac(t[l-1-i]), complement.DnaIupac(t[i])
	}
	return NewDnaIupac(string(t))
}

// Complement is the same DnaIupac with the sequence complemented
func (x *DnaIupac) Complement() (Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.DnaIupac(t[i])
	}
	return NewDnaIupac(string(t))
}
