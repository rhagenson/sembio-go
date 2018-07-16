package simple

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(Dna)

func TestDna(t *testing.T) {
	a := new(Dna)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 4))
}
