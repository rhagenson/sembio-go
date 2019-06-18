package alphabet

import "github.com/rhagenson/bio-go/bio/alphabet/internal/complement"

// Dna is the four letter standard encoding
type Dna struct {
	*Struct
}

func NewDna() *Dna {
	return &Dna{
		New(DnaLetters),
	}
}

func (*Dna) Complement(c byte) byte {
	return complement.Dna(c)
}
