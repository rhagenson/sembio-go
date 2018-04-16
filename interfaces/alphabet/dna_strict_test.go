package alphabet

import "testing"

func TestDNAStrictLength(t *testing.T) {
	var alpha = new(DNAStrict)

	if !TestAlphabetProperLength(alpha, 4) {
		t.Error("DNAStrict should have four characters.")
	}
}
