package alphabet

import "testing"

var _ Interface = new(DnaIupac)

func TestDNAIupacLength(t *testing.T) {
	if !TestAlphabetProperLength(DnaIupacLetters, 16) {
		t.Error("DNAIupac should have sixteen characters.")
	}
}
