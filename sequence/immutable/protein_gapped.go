package immutable

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/sequence"
)

var _ sequence.Reverser = new(ProteinGapped)
var _ Wither = new(ProteinGapped)

// ProteinGapped is a sequence witch validates against the ProteinGapped alphabet
// and knows how to reverse itself
type ProteinGapped struct {
	*Struct
}

// NewProteinGapped generates a New sequence that validates against the ProteinGapped alphabet
func NewProteinGapped(s string) (*ProteinGapped, error) {
	n := NewStruct(
		s,
		AlphabetIs(alphabet.ProteinGapped),
	)
	return &ProteinGapped{n}, n.Validate()
}

// Reverse is the same ProteinGapped with the sequence reversed
func (x *ProteinGapped) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewProteinGapped(string(t))
}
