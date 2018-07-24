package qscore

import (
	"math"
	"testing"
)

func TestIllumina64V13(t *testing.T) {
	want := int8(0)
	for c := minIllumina64V13; c <= maxIllumina64V13; c++ {
		got, _ := Illumina64V13(c)
		if got != want {
			t.Errorf("Illumina64V13(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestIllumina64V13Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < minIllumina64V13; c++ {
			got, err := Illumina64V13(c)
			if err == nil {
				t.Errorf("Illumina64V13(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("Illumina64V13(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(maxIllumina64V13 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := Illumina64V13(c)
			if err == nil {
				t.Errorf("Illumina64V13(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("Illumina64V13(%v) should fail, got %v value", c, got)
			}
		}
	})
}

func TestIllumina64V15(t *testing.T) {
	want := int8(2)
	for c := minIllumina64V15; c <= maxIllumina64V15; c++ {
		got, _ := Illumina64V15(c)
		if got != want {
			t.Errorf("Illumina64V15(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestIllumina64V15Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < minIllumina64V15; c++ {
			got, err := Illumina64V15(c)
			if err == nil {
				t.Errorf("Illumina64V15(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("Illumina64V15(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(maxIllumina64V15 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := Illumina64V15(c)
			if err == nil {
				t.Errorf("Illumina64V15(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("Illumina64V15(%v) should fail, got %v value", c, got)
			}
		}
	})
}
