package alphabet

import "github.com/rhagenson/bio-go/bio/alphabet/internal/complement"

// DnaIupac is the sixteen letter IUPAC encoding
type DnaIupac struct {
	*Struct
}

func NewDnaIupac() *DnaIupac {
	return &DnaIupac{
		New(DnaIupacLetters),
	}
}

func (*DnaIupac) Complement(c byte) byte {
	return complement.DnaIupac(c)
}
