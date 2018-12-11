package complement_test

import (
	"fmt"
	"testing"

	"bitbucket.org/rhagenson/bio/helpers/complement"
)

// TestIupacReturnsX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestIupacReturnsX(t *testing.T) {
	for _, c := range "XQZ" {
		usual := complement.Iupac(byte(c))
		if usual != 'X' {
			t.Errorf(
				"Want: %q, Got: %q",
				'X',
				usual,
			)
		}
	}
}

// TestIupacIsReversible checks that the complement of a complement is
// the original
func TestIupacIsReversible(t *testing.T) {
	t.Run("IUPAC is reversible", func(t *testing.T) {
		for _, c := range "RYSWKMBDHVN" {
			comp := complement.Iupac(byte(c))
			if complement.Iupac(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
			}
		}
	})
}

// BenchmarkCompIupac benchmarks the complement of each possible input byte
func BenchmarkCompIupac(b *testing.B) {
	for _, c := range "RYSWKMBDHVN" {
		b.Run(fmt.Sprintf("Complement %q", c), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				complement.Iupac(byte(c))
			}
		})
	}
}

// TestDnaIupacReturnsX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestDnaIupacReturnsX(t *testing.T) {
	for _, c := range "XQZ" {
		usual := complement.DnaIupac(byte(c))
		if usual != 'X' {
			t.Errorf(
				"Want: %q, Got: %q",
				'X',
				usual,
			)
		}
	}
}

// TestDnaIupacIsReversible checks that the complement of a complement is
// the original
func TestDnaIupacIsReversible(t *testing.T) {
	t.Run("DNA-IUPAC is reversible", func(t *testing.T) {
		for _, c := range "ATGC" + "RYSWKMBDHVN" {
			comp := complement.DnaIupac(byte(c))
			if complement.DnaIupac(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
			}
		}
	})
}

// BenchmarkCompDnaIupac benchmarks the complement of each possible input byte
func BenchmarkCompDnaIupac(b *testing.B) {
	for _, c := range "ATGC" + "RYSWKMBDHVN" {
		b.Run(fmt.Sprintf("Complement %q", c), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				complement.DnaIupac(byte(c))
			}
		})
	}
}

// TestRnaIupacReturnsX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestRnaIupacReturnsX(t *testing.T) {
	for _, c := range "XQZ" {
		usual := complement.RnaIupac(byte(c))
		if usual != 'X' {
			t.Errorf(
				"Want: %q, Got: %q",
				'X',
				usual,
			)
		}
	}
}

// TestRnaIupacIsReversible checks that the complement of a complement is
// the original
func TestRnaIupacIsReversible(t *testing.T) {
	t.Run("RNA-IUPAC is reversible", func(t *testing.T) {
		for _, c := range "AUGC" + "RYSWKMBDHVN" {
			comp := complement.RnaIupac(byte(c))
			if complement.RnaIupac(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
			}
		}
	})
}

// BenchmarkCompRnaIupac benchmarks the complement of each possible input byte
func BenchmarkCompRnaIupac(b *testing.B) {
	for _, c := range "AUGC" + "RYSWKMBDHVN" {
		b.Run(fmt.Sprintf("Complement %q", c), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				complement.RnaIupac(byte(c))
			}
		})
	}
}
