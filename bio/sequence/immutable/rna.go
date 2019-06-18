package immutable

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/sequence"
)

var _ sequence.Interface = new(Rna)
var _ sequence.Reverser = new(Rna)
var _ sequence.RevComper = new(Rna)
var _ sequence.Complementer = new(Rna)
var _ sequence.Alphabeter = new(Rna)
var _ sequence.LetterCounter = new(Rna)
var _ Wither = new(Rna)
var _ Validator = new(Rna)

// Rna is a sequence witch validates against the Rna alphabet
// and knows how to reverse, complement, and revcomp itself
type Rna struct {
	*Struct
	alpha alphabet.Interface
}

// NewRna generates a New sequence that validates against the Rna alphabet
func NewRna(s string) (*Rna, error) {
	n := New(
		s,
		AlphabetIs(alphabet.NewRna()),
	)
	return &Rna{n, alphabet.NewRna()}, n.Validate()
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
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = c.Complement(t[l-1-i]), c.Complement(t[i])
	}
	return NewRna(string(t))
}

// Complement is the same Rna with the sequence complemented
func (x *Rna) Complement() (sequence.Interface, error) {
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = c.Complement(byte(x.seq[i]))
	}
	return NewRna(string(t))
}

// Alphabet reveals the underlying alphabet in use
func (x *Rna) Alphabet() alphabet.Interface {
	return alphabet.NewRna()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *Rna) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
