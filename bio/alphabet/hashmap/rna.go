package hashmap

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/alphabet/internal/complement"
)

// Rna is the four letter standard encoding
type Rna struct {
	*Struct
}

func NewRna() *Rna {
	return &Rna{
		New(alphabet.RnaLetters),
	}
}

func (*Rna) Complement(c string) string {
	return complement.Rna(c)
}
