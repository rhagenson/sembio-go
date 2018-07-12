package helpers

// CompATGC is a bitwise manner to complement standard ATGC
// See: https://blog.kloetzl.info/reverse-complement/ for more information
// This method should be faster than CompATGCpairs
func CompATGC(c byte) byte {
	if (c & 2) > 0 {
		return c ^ 4
	}
	return c ^ 21
}

// CompATGCpairs is the usual way to complement standard ATGC
func CompATGCpairs(c byte) byte {
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
