package qscore

import (
	"fmt"
)

const (
	// minSolexa64 is the ASCII character with the lowest quality score
	minSolexa64 = byte(';')

	// maxSolexa64 is the ASCII character with the highest quality score
	maxSolexa64 = byte('h')
)

// Solexa64 takes the single-byte ASCII character used to represent
// quality score in Solexa fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an zero score and error.
func Solexa64(char byte) (int8, error) {
	if minSolexa64 <= char && char <= maxSolexa64 {
		return int8(char - minSolexa64 - 5), nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Solexa+64 symbol", char,
	)
}
