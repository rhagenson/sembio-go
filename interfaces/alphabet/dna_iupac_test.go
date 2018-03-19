package alphabet

import "testing"

func TestDNAIupacLength(t *testing.T) {
	var dna = &DNAIupac{}

	if !TestAlphabetProperLength(dna, 16) {
		t.Error("DNAIupac should only have sixteen characters")
	}
}
