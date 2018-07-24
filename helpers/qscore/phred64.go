package qscore

import (
	"fmt"
)

const (
	// minIllumina64_13 is the ASCII character with the lowest quality score
	minIllumina64V13 = byte('@')

	// maxIllumina64_13 is the ASCII character with the highest quality score
	maxIllumina64V13 = byte('h')

	// minIllumina64_15 is the ASCII character with the lowest quality score
	minIllumina64V15 = byte('B')

	// maxIllumina64_15 is the ASCII character with the highest quality score
	maxIllumina64V15 = byte('i')
)

// Illumina64V13 takes the single-byte ASCII character used to represent
// quality score in Illumina fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an zero score and error.
func Illumina64V13(char byte) (int8, error) {
	if minIllumina64V13 <= char && char <= maxIllumina64V13 {
		return int8(char - minIllumina64V13), nil
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
	if minIllumina64V15 <= char && char <= maxIllumina64V15 {
		return int8(char - minIllumina64V15 + 2), nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Illumina+64 (v1.5+) symbol", char,
	)
}
