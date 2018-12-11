package qscore_test

import (
	"math"
	"testing"

	"bitbucket.org/rhagenson/bio/helpers/qscore"
)

func TestSangerPhred33(t *testing.T) {
	want := int8(0)
	for c := qscore.MinPhred33; c <= qscore.MaxSangerPhred33; c++ {
		got, _ := qscore.SangerPhred33(c)
		if got != want {
			t.Errorf("SangerPhred33(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestSangerPhred33Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < qscore.MinPhred33; c++ {
			got, err := qscore.IlluminaPhred33(c)
			if err == nil {
				t.Errorf("SangerPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("SangerPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(qscore.MaxSangerPhred33 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := qscore.SangerPhred33(c)
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
	for c := qscore.MinPhred33; c <= qscore.MaxIlluminaPhred33; c++ {
		got, _ := qscore.IlluminaPhred33(c)
		if got != want {
			t.Errorf("IlluminaPhred33(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestIlluminaPhred33Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < qscore.MinPhred33; c++ {
			got, err := qscore.IlluminaPhred33(c)
			if err == nil {
				t.Errorf("IlluminaPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("IlluminaPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(qscore.MaxIlluminaPhred33 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := qscore.IlluminaPhred33(c)
			if err == nil {
				t.Errorf("IlluminaPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("IlluminaPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
}
