package complement

// Dna is the usual way to complement standard ATGC
func Dna(c byte) byte {
	switch c {
	case 'A':
		return 'T'
	case 'T':
		return 'A'
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	default:
		return 'X'
	}
}
