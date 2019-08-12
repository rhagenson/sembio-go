package hashmap

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/alphabet/internal/complement"
)

// Dna is the four letter standard encoding
type Dna struct {
	*Struct
}

// NewDna generates a standard DNA alphabet
func NewDna() *Dna {
	return &Dna{
		New(alphabet.DnaLetters),
	}
}

// Complement produces the standard DNA complement
func (*Dna) Complement(c string) string {
	return complement.Dna(c)
}