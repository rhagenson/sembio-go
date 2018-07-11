package illumina

import (
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
