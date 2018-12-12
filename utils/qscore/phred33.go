package qscore

import (
	"fmt"
)

const (
	// MinPhred33 is the ASCII character with the lowest quality score
	// used by both Sanger and Illumina sequences
	MinPhred33 = byte('!')

	// MaxSangerPhred33 is the ASCII character with the highest quality score
	// used by Sanger sequences
	MaxSangerPhred33 = byte('I')

	// MaxIlluminaPhred33 is the ASCII character with the highest quality score
	// used by Illumina sequences
	MaxIlluminaPhred33 = byte('J')
)

// SangerPhred33 takes the single-byte ASCII character used to represent
// quality score in Sanger fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an zero score and error.
func SangerPhred33(char byte) (int8, error) {
	if MinPhred33 <= char && char <= MaxSangerPhred33 {
		return int8(char - MinPhred33), nil
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
	if MinPhred33 <= char && char <= MaxIlluminaPhred33 {
		return int8(char - MinPhred33), nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Illumina Phred+33 (v1.8+) symbol", char,
	)
}
