package complement

import (
	"testing"
)

// TestAtgcMethodsAgree checks that the bitwise and usual way of
// generating the complement of Atgc do agree with respect to Atgc
func TestAtgcMethodsAgree(t *testing.T) {
	for _, c := range "ATGC" {
		bitwise := Atgc(byte(c))
		usual := AtgcPairs(byte(c))
		if bitwise != usual {
			t.Errorf("Bitwise comp: %q != Usual comp: %q", bitwise, usual)
		}
	}
}

// TestAtgcPairsReturnX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestAtgcPairsReturnX(t *testing.T) {
	for _, c := range "XNQZ" {
		usual := AtgcPairs(byte(c))
		if usual != 'X' {
			t.Errorf(
				"Want: %q, Got: %q",
				'X',
				usual,
			)
		}
	}
}

// TestAtgcMethodsAreReversible checks that the complement of a complement is
// the original
func TestAtgcMethodsAreReversible(t *testing.T) {
	t.Run("Pairs are reversible", func(t *testing.T) {
		for _, c := range "ATGC" {
			comp := AtgcPairs(byte(c))
			if AtgcPairs(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
			}
		}
	})
	t.Run("Bitwise method are reversible", func(t *testing.T) {
		for _, c := range "ATGC" {
			comp := Atgc(byte(c))
			if Atgc(comp) != byte(c) {
				t.Errorf("Want: %q; Got: %q", byte(c), comp)
			}
		}
	})

}

func BenchmarkAtgcMethods(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Atgc(byte('A'))
		}
	})
	b.Run("Complement T", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Atgc(byte('T'))
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

func BenchmarkAtgcPairs(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AtgcPairs(byte('A'))
		}
	})
	b.Run("Complement T", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AtgcPairs(byte('T'))
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
