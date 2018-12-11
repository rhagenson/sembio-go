package persistent

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/sequence"
)

var _ sequence.Reverser = new(Protein)
var _ Wither = new(Protein)

// Protein is a sequence witch validates against the Protein alphabet
// and knows how to reverse itself
type Protein struct {
	*Struct
}

// NewProtein generates a New sequence that validates against the Protein alphabet
func NewProtein(s string) (*Protein, error) {
	n := NewStruct(
		s,
		AlphabetIs(alphabet.Protein),
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
	return NewProtein(string(t))
}
