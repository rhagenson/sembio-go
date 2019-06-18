package mutable

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/sequence"
)

var _ sequence.Reverser = new(Protein)
var _ sequence.Alphabeter = new(Protein)
var _ sequence.LetterCounter = new(Protein)
var _ Wither = new(Protein)
var _ Validator = new(Protein)

// Protein is a sequence witch validates against the Protein alphabet
// and knows how to reverse itself
type Protein struct {
	*Struct
}

// NewProtein generates a New sequence that validates against the Protein alphabet
func NewProtein(s string) (*Protein, error) {
	n := New(
		s,
		AlphabetIs(alphabet.NewProtein()),
	)
	return &Protein{n}, n.Validate()
}

// Reverse is the same Protein with the sequence reversed
func (x *Protein) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	x.seq = string(t)
	return x, x.Validate()
}

// Alphabet reveals the underlying alphabet in use
func (x *Protein) Alphabet() alphabet.Interface {
	return alphabet.NewProtein()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *Protein) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
