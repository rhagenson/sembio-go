package alphabet

import "testing"

var _ Interface = new(RnaIupac)

func TestRNAIupacLength(t *testing.T) {
	if !TestAlphabetProperLength(RnaIupacLetters, 16) {
		t.Error("DNAIupac should have sixteen characters.")
	}
}
