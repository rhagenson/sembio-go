package sequence

// LetterCount iterates over a sequence.Interface and counts how
// many times each letter has been seen.
func LetterCount(s Interface) map[string]uint {
	c := make(map[string]uint)

	for i := uint(0); i < s.Length(); i++ {
		p, _ := s.Position(i)
		c[p]++
	}
	return c
}
