package alphabet

// TestAlphabetProperLength is a test helpher to wrap a check for
// an Alphbet implementation having the correct number of letters
func TestAlphabetProperLength(a Alphabet, l int) bool {
	return len(a.Letters()) == l
}
