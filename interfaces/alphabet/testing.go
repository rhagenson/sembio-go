package alphabet

/*
testing.go is a collection of test helpers which should wholly be used by people hoping to satisfy a
bigr interface.
*/

// TestAlphabetProperLength is a test helpher to wrap a check for
// an Alphbet implementation having the correct number of letters
func TestAlphabetProperLength(a Alphabet, l int) bool {
	return len(a.Letters()) == l
}
