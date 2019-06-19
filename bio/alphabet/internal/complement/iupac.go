package complement

// Iupac complements IUPAC ambiguous codes
func Iupac(c string) string {
	switch c {
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

// DnaIupac complements IUPAC DNA
func DnaIupac(c string) string {
	if d := Dna(c); d != "X" {
		return d
	}
	return Iupac(c)
}

// RnaIupac complements IUPAC RNA
func RnaIupac(c string) string {
	if d := Rna(c); d != "X" {
		return d
	}
	return Iupac(c)
}
