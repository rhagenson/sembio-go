package hashmap

import (
	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/internal/complement"
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
	return complement.Rna(c)
}
