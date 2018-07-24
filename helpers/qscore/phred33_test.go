package qscore

import (
	"math"
	"testing"
)

func TestSangerPhred33(t *testing.T) {
	want := int8(0)
	for c := minPhred33; c <= maxSangerPhred33; c++ {
		got, _ := SangerPhred33(c)
		if got != want {
			t.Errorf("SangerPhred33(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestSangerPhred33Errors(t *testing.T) {
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

func TestIlluminaPhred33(t *testing.T) {
	want := int8(0)
	for c := minPhred33; c <= maxIlluminaPhred33; c++ {
		got, _ := IlluminaPhred33(c)
		if got != want {
			t.Errorf("IlluminaPhred33(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestIlluminaPhred33Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < minPhred33; c++ {
			got, err := IlluminaPhred33(c)
			if err == nil {
				t.Errorf("IlluminaPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("IlluminaPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(maxIlluminaPhred33 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := IlluminaPhred33(c)
			if err == nil {
				t.Errorf("IlluminaPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("IlluminaPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
}
