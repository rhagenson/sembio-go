package complement_test

import (
	"testing"

	"bitbucket.org/rhagenson/bio/utils/complement"
)

// TestRnaReturnsX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestRnaReturnsX(t *testing.T) {
	for _, c := range "XNQZ" {
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
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			complement.Rna(byte('A'))
		}
	})
	b.Run("Complement U", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			complement.Rna(byte('U'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			complement.Rna(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			complement.Rna(byte('C'))
		}
	})
}
