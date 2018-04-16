package illumina

import (
	"testing"
)

func TestQScore(t *testing.T) {
	// Check the first 10 and last 10 valid ASCII characters for Phred+33
	// Encoding
	table := []struct {
		in  byte
		out uint8
	}{
		{'!', 0},
		{'"', 1},
		{'#', 2},
		{'$', 3},
		{'%', 4},
		{'&', 5},
		{'\'', 6},
		{'(', 7},
		{')', 8},
		{'*', 9},
		{'+', 10},
		{'@', 31},
		{'A', 32},
		{'B', 33},
		{'C', 34},
		{'D', 35},
		{'E', 36},
		{'F', 37},
		{'G', 38},
		{'H', 39},
		{'I', 40},
	}

	for _, tt := range table {
		s, _ := Phred33QScore(tt.in)
		if s != tt.out {
			t.Errorf("Calculating Q-Score with %v yielded %v, while %v was expected", tt.in, s, tt.out)
		}
	}
}
