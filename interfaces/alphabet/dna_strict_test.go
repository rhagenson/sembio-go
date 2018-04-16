package alphabet

import "testing"

func TestDNAStrictLength(t *testing.T) {
	var dna = &DNAStrict{}

	if !TestAlphabetProperLength(dna, 4) {
		t.Error("DNAStrict should have four characters.")
	}
}
