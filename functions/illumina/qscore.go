package illumina

import (
	"fmt"
)

const (
	// MaxPhredASCII is the ASCII character with the highest quality score
	MaxPhredASCII = uint8('I')

	// MinPhredASCII is the ASCII character with the lowest quality score
	MinPhredASCII = uint8('!')
)

// PhredQScore takes the single-byte ASCII character used to represent
// quality score in Illumina fastq files and returns the associated
// quality score and nil error.
// Otherwise it returns an invalid score and error.
// This is Phred+33 encoding.
func PhredQScore(r rune) (uint8, error) {
	if code := uint8(r); MinPhredASCII <= code && code <= MaxPhredASCII {
		return code - MinPhredASCII, nil
	}
	return 0, fmt.Errorf("%c is not a valid Illumina Phred quality score symbol", r)
}
