package simple

import (
	"fmt"
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

var _ alphabet.Interface = new(RnaIupac)

func TestRnaIupac(t *testing.T) {
	a := new(RnaIupac)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, '-'))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "AUGC" + "RYSWKM" + "BDHVN"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}
