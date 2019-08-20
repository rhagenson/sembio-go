package complement_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/internal/complement"
	"github.com/bio-ext/bio-go/bio/test"
)

// TestDNAReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestDNAReturnsX(t *testing.T) {
	for _, c := range alphabet.TestExcludesSingleLetters([]byte(alphabet.DnaLetters)) {
		usual := complement.Dna(string(c))
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
	t.Run("DNA is reversible", func(t *testing.T) {
		for _, c := range alphabet.DnaLetters {
			comp := complement.Dna(string(c))
			if complement.Dna(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", string(c), comp)
			}
		}
	})
}

func BenchmarkDNA(b *testing.B) {
	b.Run(fmt.Sprintf("Complement %q", alphabet.DnaLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.DnaLetters[rand.Intn(len(alphabet.DnaLetters))]
				complement.Dna(string(d))
			}
		},
	)
}
