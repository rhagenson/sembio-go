package complement

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet/simple"
)

// TestCompRnaIupacIsReversible checks that the complement of a complement is
// the original
func TestCompRnaIupacIsReversible(t *testing.T) {
	for _, c := range simple.RnaIupacLetters {
		comp := RnaIupac(byte(c))
		if RnaIupac(comp) != byte(c) {
			t.Errorf("Want: %q; Got: %q", byte(c), comp)
		}
	}
}

// TestCompDnaIupacIsReversible checks that the complement of a complement is
// the original
func TestCompDnaIupacIsReversible(t *testing.T) {
	for _, c := range simple.DnaIupacLetters {
		comp := DnaIupac(byte(c))
		if DnaIupac(comp) != byte(c) {
			t.Errorf("Want: %q; Got: %q", byte(c), comp)
		}
	}
}

func BenchmarkCompRnaIupac(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Atgc(byte('A'))
		}
	})
	b.Run("Complement U", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Atgc(byte('U'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Atgc(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Atgc(byte('C'))
		}
	})
}

func BenchmarkCompRnaIupacpairs(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AtgcPairs(byte('A'))
		}
	})
	b.Run("Complement U", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AtgcPairs(byte('U'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AtgcPairs(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AtgcPairs(byte('C'))
		}
	})
}
