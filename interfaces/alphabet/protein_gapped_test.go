package alphabet

import "testing"

func TestProteinGappedLength(t *testing.T) {
	var alpha = new(ProteinGapped)

	if !TestAlphabetProperLength(alpha, 21) {
		t.Error("ProteinGapped should have twenty-one characters.")
	}
}
