package qscore

import (
	"fmt"
)

const (
	// minPhred33 is the ASCII character with the lowest quality score
	// used by both Sanger and Illumina sequences
	minPhred33 = byte('!')

	// maxSangerPhred33 is the ASCII character with the highest quality score
	// used by Sanger sequences
	maxSangerPhred33 = byte('I')

	// maxIlluminaPhred33 is the ASCII character with the highest quality score
	// used by Illumina sequences
	maxIlluminaPhred33 = byte('J')
)

// SangerPhred33 takes the single-byte ASCII character used to represent
// quality score in Sanger fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an zero score and error.
func SangerPhred33(char byte) (int8, error) {
	if minPhred33 <= char && char <= maxSangerPhred33 {
		return int8(char - minPhred33), nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Sanger Phred+33 symbol", char,
	)
}

// IlluminaPhred33 takes the single-byte ASCII character used to represent
// quality score in Sanger fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an zero score and error.
func IlluminaPhred33(char byte) (int8, error) {
	if minPhred33 <= char && char <= maxIlluminaPhred33 {
		return int8(char - minPhred33), nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Illumina Phred+33 (v1.8+) symbol", char,
	)
}
