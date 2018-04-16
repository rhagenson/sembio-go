package alphabet

import "testing"

func TestRNAStrictLength(t *testing.T) {
	var dna = &RNAStrict{}

	if !TestAlphabetProperLength(dna, 4) {
		t.Error("RNAStrict should have four characters.")
	}
}
