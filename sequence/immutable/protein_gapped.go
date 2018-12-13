package immutable

import (
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/utils"
)

var _ sequence.Reverser = new(ProteinGapped)
var _ sequence.Alphabeter = new(ProteinGapped)
var _ sequence.LetterCounter = new(ProteinGapped)
var _ Wither = new(ProteinGapped)
var _ Validator = new(ProteinGapped)

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

// Alphabet reveals the underlying alphabet in use
func (x *ProteinGapped) Alphabet() alphabet.Interface {
	return alphabet.ProteinGapped
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *ProteinGapped) LetterCount() map[string]uint {
	return utils.LetterCount(x)
}
