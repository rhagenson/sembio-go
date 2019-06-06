package fastq_test

import (
	"math"
	"testing"

	"github.com/rhagenson/bio-go/bio/io/fastq"
)

func TestSangerPhred33(t *testing.T) {
	want := int8(0)
	for c := fastq.MinPhred33; c <= fastq.MaxSangerPhred33; c++ {
		got, _ := fastq.SangerPhred33(c)
		if got != want {
			t.Errorf("SangerPhred33(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestSangerPhred33Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < fastq.MinPhred33; c++ {
			got, err := fastq.IlluminaPhred33(c)
			if err == nil {
				t.Errorf("SangerPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("SangerPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(fastq.MaxSangerPhred33 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := fastq.SangerPhred33(c)
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
	for c := fastq.MinPhred33; c <= fastq.MaxIlluminaPhred33; c++ {
		got, _ := fastq.IlluminaPhred33(c)
		if got != want {
			t.Errorf("IlluminaPhred33(%v) got %v; want: %v", c, got, want)
		}
		want++
	}
}

func TestIlluminaPhred33Errors(t *testing.T) {
	t.Run("Below minimum errors out", func(t *testing.T) {
		for c := byte(0); c < fastq.MinPhred33; c++ {
			got, err := fastq.IlluminaPhred33(c)
			if err == nil {
				t.Errorf("IlluminaPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("IlluminaPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
	t.Run("Above maximum errors out", func(t *testing.T) {
		for c := byte(fastq.MaxIlluminaPhred33 + 1); c < byte(math.MaxUint8)/2; c++ {
			got, err := fastq.IlluminaPhred33(c)
			if err == nil {
				t.Errorf("IlluminaPhred33(%v) should fail, got nil error", c)
			}
			if got != 0 {
				t.Errorf("IlluminaPhred33(%v) should fail, got %v value", c, got)
			}
		}
	})
}
