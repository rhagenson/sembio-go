package alphabet

import "testing"

var _ Interface = new(DnaIupac)

func TestDnaIupac(t *testing.T) {
	a := new(DnaIupac)
	t.Run("Correct length", IsExpectedLength(a, 16))
	t.Run("Has gap", HasExpectedLetter(a, '-'))
}
