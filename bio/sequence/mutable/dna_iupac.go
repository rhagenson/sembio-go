package mutable

import (
	"strings"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/sequence"
)

var _ sequence.Reverser = new(DnaIupac)
var _ sequence.RevComper = new(DnaIupac)
var _ sequence.Complementer = new(DnaIupac)
var _ sequence.Alphabeter = new(DnaIupac)
var _ sequence.LetterCounter = new(DnaIupac)
var _ Wither = new(DnaIupac)

// DnaIupac is a sequence which validates against the DnaIupac alphabet
// and knows how to reverse, complement, and revcomp itself
type DnaIupac struct {
	*Struct
}

// NewDnaIupac generates a New sequence that validates against the DnaIupac alphabet
func NewDnaIupac(s string) (*DnaIupac, error) {
	n := New(
		s,
		sequence.AlphabetIs(hashmap.NewDnaIupac()),
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
	x.seq = string(t)
	return x, x.Validate()
}

// RevComp is the same DnaIupac with the sequence reversed and complemented
func (x *DnaIupac) RevComp() (sequence.Interface, error) {
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

// Complement is the same DnaIupac with the sequence complemented
func (x *DnaIupac) Complement() (sequence.Interface, error) {
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
func (x *DnaIupac) Alphabet() alphabet.Interface {
	return hashmap.NewDnaIupac()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *DnaIupac) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
