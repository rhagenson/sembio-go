package complement_test

import (
	"fmt"
	"math/rand"
	"testing"

	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/test"
	"bitbucket.org/rhagenson/bio/utils/complement"
)

// TestDNAReturnsX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestDNAReturnsX(t *testing.T) {
	for _, c := range "XNQZ" {
		usual := complement.Dna(byte(c))
		if usual != 'X' {
			t.Errorf(
				"Want: %q, Got: %q",
				'X',
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
			comp := complement.Dna(byte(c))
			if complement.Dna(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
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
				complement.Dna(d)
			}
		},
	)
}
