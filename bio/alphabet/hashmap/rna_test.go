package hashmap_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/sembio/go/bio/alphabet"
	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/test"
	"github.com/sembio/go/bio/utils"
)

func TestRna(t *testing.T) {
	var a alphabet.Interface = hashmap.NewRna()
	letters := []byte("AUGC")
	notLetters := alphabet.TestExcludesSingleLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 4))
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

// TestRnaReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestRnaReturnsX(t *testing.T) {
	a := hashmap.NewRna()
	letters := []byte("AUGC")
	notLetters := alphabet.TestExcludesSingleLetters(letters)
	for _, c := range notLetters {
		usual := a.Complement(string(c))
		if usual != "X" {
			t.Errorf(
				"Want: %q, Got: %q",
				"X",
				usual,
			)
		}
	}
}

// TestRnaIsReversible checks that the complement of a complement is
// the original
func TestRnaIsReversible(t *testing.T) {
	a := hashmap.NewRna()
	t.Run("RNA is reversible", func(t *testing.T) {
		for _, c := range "AUGC" {
			comp := a.Complement(string(c))
			if a.Complement(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", string(c), comp)
			}
		}
	})
}

func ExampleNewRna() {
	a := hashmap.NewRna()
	fmt.Println(a)
	// Output:
	// ACGU
}

func ExampleRna_Length() {
	a := hashmap.NewRna()
	fmt.Println(a.Length())
	// Output:
	// 4
}

func ExampleRna_Contains() {
	a := hashmap.NewRna()
	fmt.Println(a.Contains([]string{"A", "T", "G", "C", "U", "Q"}...))
	// Output:
	// [true false true true true false]
}

func ExampleRna_Complement() {
	a := hashmap.NewRna()
	fmt.Println(a.Complement("A"))
	// Output:
	// U
}

func BenchmarkRNA(b *testing.B) {
	a := hashmap.NewRna()
	b.Run(fmt.Sprintf("Complement %q", alphabet.RnaLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.RnaLetters[rand.Intn(len(alphabet.RnaLetters))]
				a.Complement(string(d))
			}
		},
	)
}
