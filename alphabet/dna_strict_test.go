package alphabet

import "testing"

var _ Interface = new(DNAStrict)

func TestDNAStrictLength(t *testing.T) {
	if !TestAlphabetProperLength(DnaStrictLetters, 4) {
		t.Error("DNAStrict should have four characters.")
	}
}
