package hashmap

import (
	"github.com/bio-ext/bio-go/bio/alphabet"
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
	switch c {
	case "A":
		return "T"
	case "T":
		return "A"
	case "G":
		return "C"
	case "C":
		return "G"

	case "S", "W", "N", "-":
		return c

	case "Y":
		return "R"
	case "R":
		return "Y"

	case "K":
		return "M"
	case "M":
		return "K"

	case "B":
		return "V"
	case "V":
		return "B"

	case "D":
		return "H"
	case "H":
		return "D"

	default:
		return "X"
	}
}
