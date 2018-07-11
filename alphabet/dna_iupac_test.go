package alphabet

import "testing"

var _ Interface = new(DNAIupac)

func TestDNAIupacLength(t *testing.T) {
	if !TestAlphabetProperLength(DnaIupacLetters, 16) {
		t.Error("DNAIupac should have sixteen characters.")
	}
}
