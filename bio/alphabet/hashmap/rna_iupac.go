package hashmap

import (
	"github.com/bio-ext/bio-go/bio/alphabet"
)

// RnaIupac is the sixteen letter IUPAC encoding
type RnaIupac struct {
	*Struct
}

// NewRnaIupac generates an IUPAC RNA alphabet
func NewRnaIupac() *RnaIupac {
	return &RnaIupac{
		New(alphabet.RnaIupacLetters),
	}
}

// Complement produces the IUPAC complement
func (*RnaIupac) Complement(c string) string {
	switch c {
	case "A":
		return "U"
	case "U":
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
