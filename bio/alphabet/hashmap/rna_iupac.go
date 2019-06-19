package hashmap

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/alphabet/internal/complement"
)

// RnaIupac is the sixteen letter IUPAC encoding
type RnaIupac struct {
	*Struct
}

func NewRnaIupac() *RnaIupac {
	return &RnaIupac{
		New(alphabet.RnaIupacLetters),
	}
}

func (*RnaIupac) Complement(c string) string {
	return complement.RnaIupac(c)
}
