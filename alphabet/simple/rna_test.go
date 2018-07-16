package simple

import (
	"fmt"
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(Rna)

func TestRna(t *testing.T) {
	a := new(Rna)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 4))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "AUGC"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}
