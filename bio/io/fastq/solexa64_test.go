package fastq_test

import (
	"math"
	"testing"

	"github.com/rhagenson/bio-go/bio/io/fastq"
)

func TestSolexa64(t *testing.T) {
	want := int8(-5)
	for c := fastq.MinSolexa64; c <= fastq.MaxSolexa64; c++ {
		got, _ := fastq.Solexa64(c)
		if got != want {
			t.Errorf("Solexa64(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestSolexa64Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < fastq.MinSolexa64; c++ {
			got, err := fastq.Solexa64(c)
			if err == nil {
				t.Errorf("Solexa64(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("Solexa64(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(fastq.MaxSolexa64 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := fastq.Solexa64(c)
			if err == nil {
				t.Errorf("Solexa64(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("Solexa64(%v) should fail, got %v value", c, got)
			}
		}
	})
}
