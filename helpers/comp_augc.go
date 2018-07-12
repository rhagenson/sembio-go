package helpers

// CompAUGC is a bitwise manner to complement standard AUGC
// See: https://blog.kloetzl.info/reverse-complement/ for more information
// This method should be faster than CompAUGCpairs
func CompAUGC(c byte) byte {
	if (c & 2) > 0 {
		return c ^ 4
	}
	return c ^ 20
}

// CompAUGCpairs is the usual way to complement standard AUGC
func CompAUGCpairs(c byte) byte {
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
