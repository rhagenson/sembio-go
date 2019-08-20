package immutable

import (
	"strings"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/sequence"
)

var _ sequence.Interface = new(Rna)
var _ sequence.Reverser = new(Rna)
var _ sequence.RevComper = new(Rna)
var _ sequence.Complementer = new(Rna)
var _ sequence.Alphabeter = new(Rna)
var _ sequence.LetterCounter = new(Rna)
var _ sequence.Validator = new(Rna)
var _ Wither = new(Rna)

// Rna is a sequence which validates against the Rna alphabet
// and knows how to reverse, complement, and revcomp itself
type Rna struct {
	*Struct
}

// NewRna generates a New sequence that validates against the Rna alphabet
func NewRna(s string) (*Rna, error) {
	n := New(
		s,
		sequence.AlphabetIs(hashmap.NewRna()),
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
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = byte(c.Complement(string(t[l-1-i]))[0]), byte(c.Complement(string(t[i]))[0])
	}
	return NewRna(string(t))
}

// Complement is the same Rna with the sequence complemented
func (x *Rna) Complement() (sequence.Interface, error) {
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := make([]string, l)
	var pos string
	for i := uint(0); i < l; i++ {
		pos, _ = x.Position(i)
		t[i] = c.Complement(pos)
	}
	return NewRna(strings.Join(t, ""))
}

// Alphabet reveals the underlying alphabet in use
func (x *Rna) Alphabet() alphabet.Interface {
	return hashmap.NewRna()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *Rna) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
