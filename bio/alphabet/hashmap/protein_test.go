package hashmap_test

import (
	"fmt"
	"testing"

	"github.com/sembio/go/bio/alphabet"
	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/utils"
)

func TestProtein(t *testing.T) {
	var a alphabet.Interface = hashmap.NewProtein()
	letters := []byte("ACDEFGHIKLMNPQRSTVWY")
	notLetters := alphabet.TestExcludesSingleLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 20))
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

func ExampleNewProtein() {
	a := hashmap.NewProtein()
	fmt.Println(a)
	// Output:
	// ACDEFGHIKLMNPQRSTVWY
}

func ExampleProtein_Length() {
	a := hashmap.NewProtein()
	fmt.Println(a.Length())
	// Output:
	// 20
}

func ExampleProtein_Contains() {
	a := hashmap.NewProtein()
	fmt.Println(a.Contains([]string{"A", "T", "G", "C", "U", "-", "Q"}...))
	// Output:
	// [true true true true false false true]
}
