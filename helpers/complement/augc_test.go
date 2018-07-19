package complement

import (
	"testing"
)

// TestAugcMethodsAgree checks that the bitwise and usual way of
// generating the complement of ATGC do agree with respect to ATGC
func TestAugcMethodsAgree(t *testing.T) {
	for _, c := range "AUGC" {
		bitwise := Augc(byte(c))
		usual := AugcPairs(byte(c))
		if bitwise != usual {
			t.Errorf("Bitwise comp: %q != Usual comp: %q", bitwise, usual)
		}
	}
}

// TestAugcPairsReturnX checks that when encountering an unknown nucleotide results in 'X' placeholder
func TestAugcPairsReturnX(t *testing.T) {
	for _, c := range "XNQZ" {
		usual := AugcPairs(byte(c))
		if usual != 'X' {
			t.Errorf(
				"Want: %q, Got: %q",
				'X',
				usual,
			)
		}
	}
}

func BenchmarkCompAugcMethods(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Augc(byte('A'))
		}
	})
	b.Run("Complement U", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Augc(byte('U'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Augc(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Augc(byte('C'))
		}
	})
}

func BenchmarkCompAugcPairs(b *testing.B) {
	b.Run("Complement A", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AugcPairs(byte('A'))
		}
	})
	b.Run("Complement U", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AugcPairs(byte('U'))
		}
	})
	b.Run("Complement G", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AugcPairs(byte('G'))
		}
	})
	b.Run("Complement C", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			AugcPairs(byte('C'))
		}
	})
}
