package alphabet

import "testing"

var _ Interface = new(Dna)

func TestDna(t *testing.T) {
	a := new(Dna)
	t.Run("Correct length", IsExpectedLength(a, 4))
}
