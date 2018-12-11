package complement

import (
	"testing"
)

// TestDNAReturnsX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestDNAReturnsX(t *testing.T) {
	for _, c := range "XNQZ" {
		usual := Dna(byte(c))
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
			comp := Dna(byte(c))
			if Dna(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
			}
		}
	})
}

func BenchmarkDNA(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Dna(byte('A'))
		}
	})
	b.Run("Complement T", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Dna(byte('T'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Dna(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Dna(byte('C'))
		}
	})
}
