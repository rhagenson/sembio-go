package quality

import (
	"fmt"
)

const (
	// MinSolexa64 is the ASCII character with the lowest quality score
	MinSolexa64 = byte(';')

	// MaxSolexa64 is the ASCII character with the highest quality score
	MaxSolexa64 = byte('h')
)

// Solexa64 takes the single-byte ASCII character used to represent
// quality score in Solexa qual files and returns the associated
// quality score and nil error.
// Otherwise it returns an zero score and error.
func Solexa64(char byte) (int8, error) {
	if MinSolexa64 <= char && char <= MaxSolexa64 {
		return int8(char - MinSolexa64 - 5), nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Solexa+64 symbol", char,
	)
}
