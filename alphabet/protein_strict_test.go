package alphabet

import "testing"

var _ Interface = new(ProteinStrict)

func TestProteinStrictLength(t *testing.T) {
	if !TestAlphabetProperLength(proteinStrict, 20) {
		t.Error("ProteinStrict should have twenty characters.")
	}
}
