package mutable

import (
	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/sequence"
)

var _ sequence.Reverser = new(Protein)
var _ sequence.Alphabeter = new(Protein)
var _ sequence.LetterCounter = new(Protein)
var _ sequence.Validator = new(Protein)
var _ Wither = new(Protein)

// Protein is a sequence which validates against the Protein alphabet
// and knows how to reverse itself
type Protein struct {
	*Struct
}

// NewProtein generates a New sequence that validates against the Protein alphabet
func NewProtein(s string) (*Protein, error) {
	n := New(
		s,
		sequence.AlphabetIs(hashmap.NewProtein()),
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
	return hashmap.NewProtein()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *Protein) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
