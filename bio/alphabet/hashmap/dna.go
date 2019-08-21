package hashmap

import (
	"github.com/bio-ext/bio-go/bio/alphabet"
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
	switch c {
	case "A":
		return "T"
	case "T":
		return "A"
	case "G":
		return "C"
	case "C":
		return "G"
	default:
		return "X"
	}
}
