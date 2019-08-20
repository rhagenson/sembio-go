package hashmap

import (
	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/internal/complement"
)

// DnaIupac is the sixteen letter IUPAC encoding
type DnaIupac struct {
	*Struct
}

// NewDnaIupac generates an IUPAC DNA alphabet
func NewDnaIupac() *DnaIupac {
	return &DnaIupac{
		New(alphabet.DnaIupacLetters),
	}
}

// Complement produces the IUPAC DNA complement
func (*DnaIupac) Complement(c string) string {
	return complement.DnaIupac(c)
}
