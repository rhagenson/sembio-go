package immutable

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/sequence"
	"github.com/rhagenson/bio-go/bio/utils/complement"
)

var _ sequence.Reverser = new(DnaIupac)
var _ sequence.RevComper = new(DnaIupac)
var _ sequence.Complementer = new(DnaIupac)
var _ sequence.Alphabeter = new(DnaIupac)
var _ sequence.LetterCounter = new(DnaIupac)
var _ Wither = new(DnaIupac)

// DnaIupac is a sequence witch validates against the DnaIupac alphabet
// and knows how to reverse, complement, and revcomp itself
type DnaIupac struct {
	*Struct
}

// NewDnaIupac generates a New sequence that validates against the DnaIupac alphabet
func NewDnaIupac(s string) (*DnaIupac, error) {
	n := New(
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

// Alphabet reveals the underlying alphabet in use
func (x *DnaIupac) Alphabet() alphabet.Interface {
	return alphabet.DnaIupac
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *DnaIupac) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
