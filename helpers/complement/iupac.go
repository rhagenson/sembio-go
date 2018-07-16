package complement

// IupacPairs is the usual way to
// complement IUPAC ambiguous codes via switch
func IupacPairs(c byte) byte {
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

// DnaIupac is the usual way to complement IUPAC DNA via switch
func DnaIupac(c byte) byte {
	if d := AtgcPairs(c); d != 'X' {
		return d
	}
	return IupacPairs(c)
}

// RnaIupac is the usual way to complement IUPAC RNA via switch
func RnaIupac(c byte) byte {
	if d := AugcPairs(c); d != 'X' {
		return d
	}
	return IupacPairs(c)
}
