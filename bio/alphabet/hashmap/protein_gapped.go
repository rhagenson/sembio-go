package hashmap

import "github.com/rhagenson/bio-go/bio/alphabet"

// ProteinGapped is the twenty letter standard encoding plus a gap letter
type ProteinGapped struct {
	*Struct
}

// NewProteinGapped generates a gapped protein alphabet
func NewProteinGapped() *ProteinGapped {
	return &ProteinGapped{
		New(alphabet.ProteinLetters + alphabet.GapLetter),
	}
}
