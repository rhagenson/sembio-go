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

func TestDnaIupac(t *testing.T) {
	var a alphabet.Interface = hashmap.NewDnaIupac()
	letters := []byte("ATGC" + "RYSWKM" + "BDHVN")
	notLetters := alphabet.TestExcludesSingleLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 16))
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

// TestDnaIupacReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestDnaIupacReturnsX(t *testing.T) {
	a := hashmap.NewDnaIupac()
	letters := []byte("ATGC" + "RYSWKM" + "BDHVN")
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

// TestDnaIupacIsReversible checks that the complement of a complement is
// the original
func TestDnaIupacIsReversible(t *testing.T) {
	a := hashmap.NewDnaIupac()
	t.Run("IUPAC DNA is reversible", func(t *testing.T) {
		for _, c := range alphabet.DnaIupacLetters {
			comp := a.Complement(string(c))
			if a.Complement(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", string(c), comp)
			}
		}
	})
}

func ExampleNewDnaIupac() {
	a := hashmap.NewDnaIupac()
	fmt.Println(a)
	// Output:
	// -ABCDGHKMNRSTVWY
}

func ExampleDnaIupac_Length() {
	a := hashmap.NewDnaIupac()
	fmt.Println(a.Length())
	// Output:
	// 16
}

func ExampleDnaIupac_Contains() {
	a := hashmap.NewDnaIupac()
	fmt.Println(a.Contains([]string{"A", "T", "G", "C", "-", "Q"}...))
	// Output:
	// [true true true true true false]
}

func ExampleDnaIupac_Complement() {
	a := hashmap.NewDnaIupac()
	fmt.Println(a.Complement("N"))
	// Output:
	// N
}

// BenchmarkCompDnaIupac benchmarks the complement of each possible input byte
func BenchmarkCompDnaIupac(b *testing.B) {
	a := hashmap.NewDnaIupac()
	b.Run(fmt.Sprintf("Complement %q", alphabet.DnaIupacLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.DnaIupacLetters[rand.Intn(len(alphabet.DnaIupacLetters))]
				a.Complement(string(d))
			}
		},
	)
}
