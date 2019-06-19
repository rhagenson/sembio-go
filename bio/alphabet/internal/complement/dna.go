package complement

// Dna complements standard ATGC
func Dna(c string) string {
	switch c {
	case "A":
		return "T"
	case "T":
		return "A"
	case "G":
		return "C"
	case "C":
		return "G"
	default:
		return "X"
	}
}
