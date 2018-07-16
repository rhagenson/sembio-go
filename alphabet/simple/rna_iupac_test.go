package simple

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(RnaIupac)

func TestRnaIupac(t *testing.T) {
	a := new(RnaIupac)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, '-'))
}
