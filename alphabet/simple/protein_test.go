package simple

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(Protein)

func TestProtein(t *testing.T) {
	a := new(Protein)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 20))
}
