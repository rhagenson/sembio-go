package hashmap

import (
	"github.com/sembio/go/bio/alphabet"
)

// Rna is the four letter standard encoding
type Rna struct {
	*Struct
}

// NewRna generates a standard RNA alphabet
func NewRna() *Rna {
	return &Rna{
		New(alphabet.RnaLetters),
	}
}

// Complement produces the standard RNA complement
func (*Rna) Complement(c string) string {
	switch c {
	case "A":
		return "U"
	case "U":
		return "A"
	case "G":
		return "C"
	case "C":
		return "G"
	default:
		return "X"
	}
}
