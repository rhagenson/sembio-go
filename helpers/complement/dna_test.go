package complement_test

import (
	"testing"

	"bitbucket.org/rhagenson/bio/helpers/complement"
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
		for _, c := range "ATGC" {
			comp := complement.Dna(byte(c))
			if complement.Dna(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
			}
		}
	})
}

func BenchmarkDNA(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			complement.Dna(byte('A'))
		}
	})
	b.Run("Complement T", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			complement.Dna(byte('T'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			complement.Dna(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			complement.Dna(byte('C'))
		}
	})
}
