package complement_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/internal/complement"
	"github.com/bio-ext/bio-go/bio/test"
)

// TestIupacReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestIupacReturnsX(t *testing.T) {
	for _, c := range string(alphabet.TestExcludesSingleLetters([]byte(alphabet.IupacLetters))) {
		usual := complement.Iupac(string(c))
		if usual != "X" {
			t.Errorf(
				"Want: %q, Got: %q",
				"X",
				usual,
			)
		}
	}
}

// TestIupacIsReversible checks that the complement of a complement is
// the original
func TestIupacIsReversible(t *testing.T) {
	t.Run("IUPAC is reversible", func(t *testing.T) {
		for _, c := range alphabet.IupacLetters {
			comp := complement.Iupac(string(c))
			if complement.Iupac(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
			}
		}
	})
}

// BenchmarkCompIupac benchmarks the complement of each possible input byte
func BenchmarkCompIupac(b *testing.B) {
	b.Run(fmt.Sprintf("Complement %q", alphabet.IupacLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.IupacLetters[rand.Intn(len(alphabet.IupacLetters))]
				complement.Iupac(string(d))
			}
		},
	)
}

// TestDnaIupacReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestDnaIupacReturnsX(t *testing.T) {
	for _, c := range "XQZ" {
		usual := complement.DnaIupac(string(c))
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
	t.Run("DNA-IUPAC is reversible", func(t *testing.T) {
		for _, c := range alphabet.DnaIupacLetters {
			comp := complement.DnaIupac(string(c))
			if complement.DnaIupac(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", string(c), comp)
			}
		}
	})
}

// BenchmarkCompDnaIupac benchmarks the complement of each possible input byte
func BenchmarkCompDnaIupac(b *testing.B) {
	b.Run(fmt.Sprintf("Complement %q", alphabet.DnaIupacLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.DnaIupacLetters[rand.Intn(len(alphabet.DnaIupacLetters))]
				complement.DnaIupac(string(d))
			}
		},
	)
}

// TestRnaIupacReturnsX checks that when encountering an unknown nucleotide results in "X" placeholder
func TestRnaIupacReturnsX(t *testing.T) {
	for _, c := range "XQZ" {
		usual := complement.RnaIupac(string(c))
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
	t.Run("RNA-IUPAC is reversible", func(t *testing.T) {
		for _, c := range alphabet.RnaIupacLetters {
			comp := complement.RnaIupac(string(c))
			if complement.RnaIupac(comp) != string(c) {
				t.Errorf("Want: %q; Got: %q", string(c), comp)
			}
		}
	})
}

// BenchmarkCompRnaIupac benchmarks the complement of each possible input byte
func BenchmarkCompRnaIupac(b *testing.B) {
	b.Run(fmt.Sprintf("Complement %q", alphabet.RnaIupacLetters),
		func(b *testing.B) {
			rand.Seed(test.Seed)
			var d byte
			for n := 0; n < b.N; n++ {
				d = alphabet.RnaIupacLetters[rand.Intn(len(alphabet.RnaIupacLetters))]
				complement.RnaIupac(string(d))
			}
		},
	)
}
