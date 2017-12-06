package functions

import "fmt"

const (
	MAX_Q_SCORE   = 40
	MIN_Q_SCORE   = 0
	MAX_ASCII_VAL = 73
	MIN_ASCII_VAL = 33
)

// IlluminaQScore takes the single-byte ASCII character used to represent
// quality score in Illumina fastq files and returns the associated
// quality score. Otherwise it returns a  quality score and error
func IlluminaQScore(r rune) (q uint8, e err) {
	if code := uint8(r); MINI_ASCII_VAL < code {
		if q = code - MIN_ASCII_VAL; q < MAX_Q_SCORE {
			return
		} else {
			e = fmt.Errorf("%v is not a valid Illumina quality score symbol.", r)
			return
		}
	}
}
