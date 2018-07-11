package alphabet

import "testing"

var _ Interface = new(RNAStrict)

func TestRNAStrictLength(t *testing.T) {
	if !TestAlphabetProperLength(RnaStrictLetters, 4) {
		t.Error("RNAStrict should have four characters.")
	}
}
