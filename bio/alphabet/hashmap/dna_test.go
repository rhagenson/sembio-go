package hashmap_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/test"
	"github.com/bio-ext/bio-go/bio/utils"
)

func TestDna(t *testing.T) {
	var a alphabet.Interface = hashmap.NewDna()
	letters := []byte("ATGC")
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

// TestDNAReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestDNAReturnsX(t *testing.T) {
	a := hashmap.NewDna()
	letters := []byte("ATGC")
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

// TestDnaIsReversible checks that the complement of a complement is
// the original
func TestDnaIsReversible(t *testing.T) {
	a := hashmap.NewDna()
	t.Run("DNA is reversible", func(t *testing.T) {
		for _, c := range alphabet.DnaLetters {
			comp := a.Complement(string(c))
			if a.Complement(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", string(c), comp)
			}
		}
	})
}

func ExampleNewDna() {
	a := hashmap.NewDna()
	fmt.Println(a)
	// Output:
	// ACGT
}

func ExampleDna_Length() {
	a := hashmap.NewDna()
	fmt.Println(a.Length())
	// Output:
	// 4
}

func ExampleDna_Contains() {
	a := hashmap.NewDna()
	fmt.Println(a.Contains([]string{"A", "T", "G", "C", "U", "Q"}...))
	// Output:
	// [true true true true false false]
}

func ExampleDna_Complement() {
	a := hashmap.NewDna()
	fmt.Println(a.Complement("A"))
	// Output:
	// T
}

func BenchmarkDNA(b *testing.B) {
	a := hashmap.NewDna()
	b.Run(fmt.Sprintf("Complement %q", alphabet.DnaLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.DnaLetters[rand.Intn(len(alphabet.DnaLetters))]
				a.Complement(string(d))
			}
		},
	)
}
