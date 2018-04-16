package alphabet

import "testing"

func TestDNAIupacLength(t *testing.T) {
	var alpha = new(DNAIupac)

	if !TestAlphabetProperLength(alpha, 16) {
		t.Error("DNAIupac should have sixteen characters.")
	}
}
