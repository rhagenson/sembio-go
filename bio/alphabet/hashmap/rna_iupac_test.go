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

func TestRnaIupac(t *testing.T) {
	var a alphabet.Interface = hashmap.NewRnaIupac()
	letters := []byte("AUGC" + "RYSWKM" + "BDHVN")
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

// TestRnaIupacReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestRnaIupacReturnsX(t *testing.T) {
	a := hashmap.NewRnaIupac()
	letters := []byte("AUGC" + "RYSWKM" + "BDHVN")
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

// TestRnaIupacIsReversible checks that the complement of a complement is
// the original
func TestRnaIupacIsReversible(t *testing.T) {
	a := hashmap.NewRnaIupac()
	t.Run("IUPAC RNA is reversible", func(t *testing.T) {
		for _, c := range alphabet.RnaIupacLetters {
			comp := a.Complement(string(c))
			if a.Complement(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", string(c), comp)
			}
		}
	})
}

func ExampleNewRnaIupac() {
	a := hashmap.NewRnaIupac()
	fmt.Println(a)
	// Output:
	// -ABCDGHKMNRSUVWY
}

func ExampleRnaIupac_Length() {
	a := hashmap.NewRnaIupac()
	fmt.Println(a.Length())
	// Output:
	// 16
}

func ExampleRnaIupac_Contains() {
	a := hashmap.NewRnaIupac()
	fmt.Println(a.Contains([]string{"A", "T", "G", "C", "U", "Q"}...))
	// Output:
	// [true false true true true false]
}

func ExampleRnaIupac_Complement() {
	a := hashmap.NewRnaIupac()
	fmt.Println(a.Complement("N"))
	// Output:
	// N
}

// BenchmarkCompRnaIupac benchmarks the complement of each possible input byte
func BenchmarkCompRnaIupac(b *testing.B) {
	a := hashmap.NewRnaIupac()
	b.Run(fmt.Sprintf("Complement %q", alphabet.RnaIupacLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.RnaIupacLetters[rand.Intn(len(alphabet.RnaIupacLetters))]
				a.Complement(string(d))
			}
		},
	)
}
