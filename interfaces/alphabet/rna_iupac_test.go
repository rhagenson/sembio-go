package alphabet

import "testing"

func TestRNAIupacLength(t *testing.T) {
	var alpha = new(RNAIupac)

	if !TestAlphabetProperLength(alpha, 16) {
		t.Error("DNAIupac should have sixteen characters.")
	}
}
