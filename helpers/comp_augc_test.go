package helpers

import (
	"testing"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

// TestCompAUGCMethodsAgree checks that the bitwise and usual way of
// generating the complement of ATGC do agree with respect to ATGC
func TestCompAUGCMethodsAgree(t *testing.T) {
	for _, c := range alphabet.RnaLetters {
		bitwise := CompAUGC(byte(c))
		usual := CompAUGCpairs(byte(c))
		if bitwise != usual {
			t.Errorf("Bitwise comp: %q != Usual comp: %q", bitwise, usual)
		}
	}
}

func BenchmarkCompAUGC(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGC(byte('A'))
		}
	})
	b.Run("Complement U", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGC(byte('U'))
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

func BenchmarkCompAUGCpairs(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGCpairs(byte('A'))
		}
	})
	b.Run("Complement U", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			CompATGCpairs(byte('U'))
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
