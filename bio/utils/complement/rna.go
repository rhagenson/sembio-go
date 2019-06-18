package complement

// Rna complements standard AUGC
func Rna(c byte) byte {
	switch c {
	case 'A':
		return 'U'
	case 'U':
		return 'A'
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	default:
		return 'X'
	}
}