package illumina

import (
	"fmt"
)

const (
	// Phred33Max is the ASCII character with the highest quality score
	Phred33Max = byte('I')

	// Phred33Min is the ASCII character with the lowest quality score
	Phred33Min = byte('!')
)

// Phred33QScore takes the single-byte ASCII character used to represent
// quality score in Illumina fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an invalid score and error.
func Phred33QScore(char byte) (uint8, error) {
	if Phred33Min <= char && char <= Phred33Max {
		return char - Phred33Min, nil
	}
	return 0, fmt.Errorf("%q is not a valid Illumina Phred quality score symbol", char)
}
