package alphabet

import "testing"

var _ Interface = new(Rna)

func TestRna(t *testing.T) {
	a := new(Rna)
	t.Run("Correct length", IsExpectedLength(a, 4))
}
