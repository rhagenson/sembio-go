package alphabet

import "testing"

func TestProteinStrictLength(t *testing.T) {
	var alpha = new(ProteinStrict)

	if !TestAlphabetProperLength(alpha, 20) {
		t.Error("ProteinStrict should have twenty characters.")
	}
}
