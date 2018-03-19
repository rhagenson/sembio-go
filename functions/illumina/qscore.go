package illumina

import (
	"fmt"
)

const (
	// MaxASCIIValue is the ASCII character with the highest quality score
	MaxASCIIValue = uint8('I')

	// MinASCIIValue is the ASCII character with the lowest quality score
	MinASCIIValue = uint8('!')
)

// QScore takes the single-byte ASCII character used to represent
// quality score in Illumina fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an invalid score and error.
// This is Phred+33 encoding.
func QScore(r rune) (uint8, error) {
	if code := uint8(r); MinASCIIValue <= code && code <= MaxASCIIValue {
		return code - MinASCIIValue, nil
	}
	return 0, fmt.Errorf("%v is not a valid Illumina quality score symbol", r)
}
