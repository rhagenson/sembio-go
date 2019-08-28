package fastq

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/bio-ext/bio-go/bio/sequence"
)

// Read reads n records from a FASTQ file using the generator f to validate the sequences
// Only records up to the first error are returned (along with the error)
func Read(r io.Reader, n uint, f sequence.Generator) ([]Interface, error) {
	br := bufio.NewScanner(r)
	br.Split(bufio.ScanLines)
	records := make([]Interface, 0, n)
	count := uint(0)
	header := ""
	plusLine := ""
	seq := new(strings.Builder)
	quality := new(strings.Builder)
	for br.Scan() {
		line := strings.TrimSpace(br.Text())
		if line != "" && line[0] == FastqHeaderPrefix { // New record
			// Line 1: Header
			header = line[1:]

			// Line2: Sequence
			br.Scan()
			seq.WriteString(strings.TrimSpace(br.Text()))
			seqx, err := f(seq.String())
			if err != nil {
				return records, err
			}

			// Line 3: Header (should match Line 1 above)
			br.Scan()
			plusLine = strings.TrimSpace(br.Text())
			if plusLine[0] == FastqPreQualityHeaderPrefix {
				if plusLine[1:] != header {
					return records, fmt.Errorf("first header:\n\t%q\ndid not match second header:\n\t%q",
						header, strings.TrimSpace(br.Text()))
				}
			} else {
				return records, fmt.Errorf("second header line did not start with %q", FastqPreQualityHeaderPrefix)
			}

			// Line 4: Quality
			br.Scan()
			quality.WriteString(strings.TrimSpace(br.Text()))

			// Add record
			records = append(records, New(header, quality.String(), seqx))
			count++
			if n != 0 && count == n {
				return records, nil
			}
		}
	}
	if header != "" && seq.String() != "" && quality.String() != "" {
		seqx, err := f(seq.String())
		if err != nil {
			return records, err
		}
		records = append(records, New(header, quality.String(), seqx))
	}
	return records, nil
}

// ReadSingle reads a single records from a FASTQ file using the generator f to validate the sequence
func ReadSingle(r io.Reader, f sequence.Generator) (Interface, error) {
	records, err := Read(r, 1, f)
	return records[0], err
}

// ReadMulti reads all records from a FASTQ file using the generator f to validate the sequences
// Only records up to the first error are returned (along with the error)
func ReadMulti(r io.Reader, f sequence.Generator) ([]Interface, error) {
	return Read(r, 0, f)
}
