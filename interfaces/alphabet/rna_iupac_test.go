package alphabet

import "testing"

func TestRNAIupacLength(t *testing.T) {
	var rna = &RNAIupac{}

	if !TestAlphabetProperLength(rna, 16) {
		t.Error("DNAIupac should have sixteen characters.")
	}
}
