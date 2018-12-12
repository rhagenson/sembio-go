package qscore

import (
	"fmt"
)

const (
	// MinIllumina64V13 is the ASCII character with the lowest quality score
	MinIllumina64V13 = byte('@')

	// MaxIllumina64V13 is the ASCII character with the highest quality score
	MaxIllumina64V13 = byte('h')

	// MinIllumina64V15 is the ASCII character with the lowest quality score
	MinIllumina64V15 = byte('B')

	// MaxIllumina64V15 is the ASCII character with the highest quality score
	MaxIllumina64V15 = byte('i')
)

// Illumina64V13 takes the single-byte ASCII character used to represent
// quality score in Illumina fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an zero score and error.
func Illumina64V13(char byte) (int8, error) {
	if MinIllumina64V13 <= char && char <= MaxIllumina64V13 {
		return int8(char - MinIllumina64V13), nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Illumina+64 (v1.3+) symbol", char,
	)
}

// Illumina64V15 takes the single-byte ASCII character used to represent
// quality score in Illumina fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an zero score and error.
// Illumina+64 (v1.5) handles the values [0,2] the following way:
//		0: unused
//		1: unused
//		2: Read Segment Quality Control Indicator
func Illumina64V15(char byte) (int8, error) {
	if MinIllumina64V15 <= char && char <= MaxIllumina64V15 {
		return int8(char - MinIllumina64V15 + 2), nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Illumina+64 (v1.5+) symbol", char,
	)
}
