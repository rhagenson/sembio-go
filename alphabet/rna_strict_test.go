package alphabet

import "testing"

var _ Interface = new(RnaStrict)

func TestRNAStrictLength(t *testing.T) {
	if !TestAlphabetProperLength(RnaStrictLetters, 4) {
		t.Error("RNAStrict should have four characters.")
	}
}
