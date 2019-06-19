package mutable

import (
	"strings"

	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/alphabet/hashmap"
	"github.com/rhagenson/bio-go/bio/sequence"
)

var _ sequence.Reverser = new(RnaIupac)
var _ sequence.RevComper = new(RnaIupac)
var _ sequence.Complementer = new(RnaIupac)
var _ sequence.Alphabeter = new(RnaIupac)
var _ sequence.LetterCounter = new(RnaIupac)
var _ sequence.Validator = new(RnaIupac)
var _ Wither = new(RnaIupac)

// RnaIupac is a sequence witch validates against the RnaIupac alphabet
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
	x.seq = string(t)
	return x, x.Validate()
}

// RevComp is the same RnaIupac with the sequence reversed and complemented
func (x *RnaIupac) RevComp() (sequence.Interface, error) {
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := make([]string, l)
	var pos1, pos2 string
	for i := uint(0); i < l/2; i++ {
		pos1, _ = x.Position(i)
		pos2, _ = x.Position(l - 1 - i)
		t[i], t[l-1-i] = c.Complement(pos2), c.Complement(pos1)
	}
	x.seq = strings.Join(t, "")
	return x, x.Validate()
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
	x.seq = strings.Join(t, "")
	return x, x.Validate()
}

// Alphabet reveals the underlying alphabet in use
func (x *RnaIupac) Alphabet() alphabet.Interface {
	return hashmap.NewRnaIupac()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *RnaIupac) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
