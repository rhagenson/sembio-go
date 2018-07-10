package alphabet

import "testing"

var _ Interface = new(ProteinGapped)

func TestProteinGappedLength(t *testing.T) {
	if !TestAlphabetProperLength(proteinGapped, 21) {
		t.Error("ProteinGapped should have twenty-one characters.")
	}
}
