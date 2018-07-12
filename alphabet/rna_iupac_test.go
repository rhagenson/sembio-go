package alphabet

import "testing"

var _ Interface = new(RnaIupac)

func TestRnaIupac(t *testing.T) {
	a := new(RnaIupac)
	t.Run("Correct length", IsExpectedLength(a, 16))
	t.Run("Has gap", HasExpectedLetter(a, '-'))
}
