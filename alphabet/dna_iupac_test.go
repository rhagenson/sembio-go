package alphabet

import "testing"

var _ Interface = new(DNAIupac)

func TestDNAIupacLength(t *testing.T) {
	if !TestAlphabetProperLength(dnaIupacLetters, 16) {
		t.Error("DNAIupac should have sixteen characters.")
	}
}
