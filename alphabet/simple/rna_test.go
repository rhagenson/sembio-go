package simple

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(Rna)

func TestRna(t *testing.T) {
	a := new(Rna)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 4))
}
