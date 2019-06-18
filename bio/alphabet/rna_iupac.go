package alphabet

import "github.com/rhagenson/bio-go/bio/alphabet/internal/complement"

// RnaIupac is the sixteen letter IUPAC encoding
type RnaIupac struct {
	*Struct
}

func NewRnaIupac() *RnaIupac {
	return &RnaIupac{
		New(RnaIupacLetters),
	}
}

func (*RnaIupac) Complement(c byte) byte {
	return complement.RnaIupac(c)
}
