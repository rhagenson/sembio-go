package complement

import (
	"testing"
)

// TestRnaIupacIsReversible checks that the complement of a complement is
// the original
func TestRnaIupacIsReversible(t *testing.T) {
	for _, c := range "AUGC" + "RYSWKM" + "BDHVN" {
		comp := RnaIupac(byte(c))
		if RnaIupac(comp) != byte(c) {
			t.Errorf("Want: %q; Got: %q", byte(c), comp)
		}
	}
}

// TestIupacPairsReturnX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestIupacPairsReturnX(t *testing.T) {
	for _, c := range "XQZ" {
		usual := IupacPairs(byte(c))
		if usual != 'X' {
			t.Errorf(
				"Want: %q, Got: %q",
				'X',
				usual,
			)
		}
	}
}

// TestCompDnaIupacIsReversible checks that the complement of a complement is
// the original
func TestCompDnaIupacIsReversible(t *testing.T) {
	for _, c := range "ATGC" + "RYSWKM" + "BDHVN" {
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
