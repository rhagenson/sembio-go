package complement_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/test"
	"github.com/rhagenson/bio-go/bio/utils/complement"
)

// TestRnaReturnsX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestRnaReturnsX(t *testing.T) {
	for _, c := range string(alphabet.TestExcludesLetters([]byte(alphabet.RnaLetters))) {
		usual := complement.Rna(byte(c))
		if usual != 'X' {
			t.Errorf(
				"Want: %q, Got: %q",
				'X',
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
			comp := complement.Rna(byte(c))
			if complement.Rna(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
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
				complement.Rna(d)
			}
		},
	)
}
