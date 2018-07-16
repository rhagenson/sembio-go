package simple

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(DnaIupac)

func TestDnaIupac(t *testing.T) {
	a := new(DnaIupac)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, '-'))
}
