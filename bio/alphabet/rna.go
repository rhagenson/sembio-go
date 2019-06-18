package alphabet

import "github.com/rhagenson/bio-go/bio/alphabet/internal/complement"

// Rna is the four letter standard encoding
type Rna struct {
	*Struct
}

func NewRna() *Rna {
	return &Rna{
		New(RnaLetters),
	}
}

func (*Rna) Complement(c byte) byte {
	return complement.Rna(c)
}
