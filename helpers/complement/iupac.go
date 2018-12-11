package complement

// Iupac is the usual way to
// complement IUPAC ambiguous codes
func Iupac(c byte) byte {
	switch c {
	case 'S', 'W', 'N', '-':
		return c

	case 'Y':
		return 'R'
	case 'R':
		return 'Y'

	case 'K':
		return 'M'
	case 'M':
		return 'K'

	case 'B':
		return 'V'
	case 'V':
		return 'B'

	case 'D':
		return 'H'
	case 'H':
		return 'D'

	default:
		return 'X'
	}
}

// DnaIupac is the usual way to complement IUPAC DNA
func DnaIupac(c byte) byte {
	if d := Dna(c); d != 'X' {
		return d
	}
	return Iupac(c)
}

// RnaIupac is the usual way to complement IUPAC RNA
func RnaIupac(c byte) byte {
	if d := Rna(c); d != 'X' {
		return d
	}
	return Iupac(c)
}
