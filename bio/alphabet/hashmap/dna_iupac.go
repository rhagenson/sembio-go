package hashmap

import (
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/alphabet/internal/complement"
)

// DnaIupac is the sixteen letter IUPAC encoding
type DnaIupac struct {
	*Struct
}

func NewDnaIupac() *DnaIupac {
	return &DnaIupac{
		New(alphabet.DnaIupacLetters),
	}
}

func (*DnaIupac) Complement(c string) string {
	return complement.DnaIupac(c)
}
