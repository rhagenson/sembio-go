package alphabet

import "testing"

func TestRNAStrictLength(t *testing.T) {
	var alpha = new(RNAStrict)

	if !TestAlphabetProperLength(alpha, 4) {
		t.Error("RNAStrict should have four characters.")
	}
}
