package helpers

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

// TestCompATGCMethodsAgree checks that the bitwise and usual way of
// generating the complement of ATGC do agree with respect to ATGC
func TestCompATGCMethodsAgree(t *testing.T) {
	for _, c := range alphabet.DnaStrictLetters {
		bitwise := CompATGC(byte(c))
		usual := CompATGCpairs(byte(c))
		if bitwise != usual {
			t.Errorf("Bitwise comp: %q != Usual comp: %q", bitwise, usual)
		}
	}
}

func BenchmarkCompATGC(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGC(byte('A'))
		}
	})
	b.Run("Complement T", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGC(byte('T'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGC(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGC(byte('C'))
		}
	})
}

func BenchmarkCompATGCpairs(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGCpairs(byte('A'))
		}
	})
	b.Run("Complement T", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGCpairs(byte('T'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGCpairs(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGCpairs(byte('C'))
		}
	})
}
