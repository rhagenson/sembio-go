package illumina

import (
	"math"
	"testing"
)

func TestPhred33QScore(t *testing.T) {
	want := uint8(0)
	for c := MinPhred33; c <= MaxPhred33; c++ {
		got, _ := Phred33QScore(c)
		if got != want {
			t.Errorf("Phred33QScore(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestPhred33QScoreErrors(t *testing.T) {
	t.Run("Below MinPhred33 errors out", func(t *testing.T) {
		for c := byte(0); c < MinPhred33; c++ {
			got, err := Phred33QScore(c)
			if err == nil {
				t.Errorf("Phred33QScore(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("Phred33QScore(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above MaxPhred33 errors out", func(t *testing.T) {
		for c := byte(MaxPhred33 + 1); c < byte(math.MaxUint8); c++ {
			got, err := Phred33QScore(c)
			if err == nil {
				t.Errorf("Phred33QScore(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("Phred33QScore(%v) should fail, got %v value", c, got)
			}
		}
	})
}
