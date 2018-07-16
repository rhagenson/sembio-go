package simple

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(ProteinGapped)

func TestProteinGapped(t *testing.T) {
	a := new(ProteinGapped)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 21))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, '-'))
}
