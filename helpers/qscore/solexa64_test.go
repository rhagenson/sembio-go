package qscore

import (
	"math"
	"testing"
)

func TestSolexa64(t *testing.T) {
	want := int8(-5)
	for c := minSolexa64; c <= maxSolexa64; c++ {
		got, _ := Solexa64(c)
		if got != want {
			t.Errorf("Solexa64(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestSolexa64Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < minPhred33; c++ {
			got, err := IlluminaPhred33(c)
			if err == nil {
				t.Errorf("SangerPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("SangerPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(maxSangerPhred33 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := SangerPhred33(c)
			if err == nil {
				t.Errorf("SangerPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("SangerPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
}
