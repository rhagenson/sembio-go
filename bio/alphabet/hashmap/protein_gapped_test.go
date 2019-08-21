package hashmap_test

import (
	"fmt"
	"testing"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/utils"
)

func TestProteinGapped(t *testing.T) {
	var a alphabet.Interface = hashmap.NewProteinGapped()
	letters := []byte("ACDEFGHIKLMNPQRSTVWY")
	notLetters := alphabet.TestExcludesSingleLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 21))
	t.Run("Has gap", alphabet.TestHasExpectedLetter(a, "-"))
	t.Run("Expected letters", func(t *testing.T) {
		for i, v := range a.Contains(utils.BytesToStrings(letters)...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		for i, v := range a.Contains(utils.BytesToStrings(notLetters)...) {
			t.Run(fmt.Sprintf("Excludes %q", notLetters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", notLetters[i])
				}
			})
		}
	})
}

func ExampleNewProteinGapped() {
	a := hashmap.NewProteinGapped()
	fmt.Println(a)
	// Output:
	// -ACDEFGHIKLMNPQRSTVWY
}

func ExampleProteinGapped_Length() {
	a := hashmap.NewProteinGapped()
	fmt.Println(a.Length())
	// Output:
	// 21
}

func ExampleProteinGapped_Contains() {
	a := hashmap.NewProteinGapped()
	fmt.Println(a.Contains([]string{"A", "T", "G", "C", "U", "-", "Q"}...))
	// Output:
	// [true true true true false true true]
}
