package complement_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/internal/complement"
	"github.com/bio-ext/bio-go/bio/test"
)

// TestRnaReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestRnaReturnsX(t *testing.T) {
	for _, c := range string(alphabet.TestExcludesSingleLetters([]byte(alphabet.RnaLetters))) {
		usual := complement.Rna(string(c))
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
	t.Run("RNA is reversible", func(t *testing.T) {
		for _, c := range "AUGC" {
			comp := complement.Rna(string(c))
			if complement.Rna(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", string(c), comp)
			}
		}
	})
}

func BenchmarkRNA(b *testing.B) {
	b.Run(fmt.Sprintf("Complement %q", alphabet.RnaLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.RnaLetters[rand.Intn(len(alphabet.RnaLetters))]
				complement.Rna(string(d))
			}
		},
	)
}
