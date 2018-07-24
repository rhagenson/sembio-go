package illumina

import (
	"fmt"
)

const (
	// MaxPhred33 is the ASCII character with the highest quality score
	MaxPhred33 = byte('I')

	// MinPhred33 is the ASCII character with the lowest quality score
	MinPhred33 = byte('!')
)

// Phred33QScore takes the single-byte ASCII character used to represent
// quality score in Illumina fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an invalid score and error.
func Phred33QScore(char byte) (uint8, error) {
	if MinPhred33 <= char && char <= MaxPhred33 {
		return char - MinPhred33, nil
	}
	return 0, fmt.Errorf(
		"%q is not a valid Illumina Phred33 symbol", char,
	)
}
