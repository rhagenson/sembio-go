package simple

import (
	"fmt"
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(ProteinGapped)

func TestProteinGapped(t *testing.T) {
	a := new(ProteinGapped)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 21))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, '-'))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "ACDEFGHIKLMNPQRSTVWY"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}
