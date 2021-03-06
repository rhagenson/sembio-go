package immutable

import (
	"strings"

	"github.com/sembio/go/bio/alphabet"
	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/sequence"
)

var _ sequence.Reverser = new(RnaIupac)
var _ sequence.RevComper = new(RnaIupac)
var _ sequence.Complementer = new(RnaIupac)
var _ sequence.Alphabeter = new(RnaIupac)
var _ sequence.LetterCounter = new(RnaIupac)
var _ sequence.Validator = new(RnaIupac)
var _ Wither = new(RnaIupac)

// RnaIupac is a sequence which validates against the RnaIupac alphabet
// and knows how to reverse, complement, and revcomp itself
type RnaIupac struct {
	*Struct
}

// NewRnaIupac generates a New sequence that validates against the RnaIupac alphabet
func NewRnaIupac(s string) (*RnaIupac, error) {
	n := New(
		s,
		sequence.AlphabetIs(hashmap.NewRnaIupac()),
	)
	return &RnaIupac{n}, n.Validate()
}

// Reverse is the same RnaIupac with the sequence reversed
func (x *RnaIupac) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewRnaIupac(string(t))
}

// RevComp is the same RnaIupac with the sequence reversed and complemented
func (x *RnaIupac) RevComp() (sequence.Interface, error) {
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = byte(c.Complement(string(t[l-1-i]))[0]), byte(c.Complement(string(t[i]))[0])
	}
	return NewRnaIupac(string(t))
}

// Complement is the same RnaIupac with the sequence complemented
func (x *RnaIupac) Complement() (sequence.Interface, error) {
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := make([]string, l)
	var pos string
	for i := uint(0); i < l; i++ {
		pos, _ = x.Position(i)
		t[i] = c.Complement(pos)
	}
	return NewRnaIupac(strings.Join(t, ""))
}

// Alphabet reveals the underlying alphabet in use
func (x *RnaIupac) Alphabet() alphabet.Interface {
	return hashmap.NewRnaIupac()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *RnaIupac) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
