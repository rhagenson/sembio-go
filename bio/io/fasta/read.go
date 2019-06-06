package fasta

import (
	"bufio"
	"io"
	"strings"

	"github.com/rhagenson/bio-go/bio/sequence"
)

// Read reads n records from a FASTA file using the generator f to validate the sequences
// Only records up to the first error are returned (along with the error)
func Read(r io.Reader, n uint, f sequence.Generator) ([]Interface, error) {
	br := bufio.NewScanner(r)
	br.Split(bufio.ScanLines)
	records := make([]Interface, 0, n)
	count := uint(0)
	header := ""
	seq := new(strings.Builder)
	for br.Scan() {
		line := strings.TrimSpace(br.Text())
		switch {
		case line == "": // Skip blank lines
			continue
		case line[0] == FastaHeaderPrefix: // Header
			if seq.Len() == 0 {
				header = line
			} else {
				seqx, err := f(seq.String())
				if err != nil {
					return records, err
				}
				records = append(records, New(header, seqx))
				count++
				if n != 0 && count == n {
					return records, nil
				}
				header = ""
				seq = new(strings.Builder)
			}
		default: // Sequence
			seq.WriteString(line)
		}
	}
	if header != "" && seq.String() != "" {
		seqx, err := f(seq.String())
		if err != nil {
			return records, err
		}
		records = append(records, New(header, seqx))
	}
	return records, nil
}

// ReadSingle reads a single records from a FASTA file using the generator f to validate the sequence
func ReadSingle(r io.Reader, f sequence.Generator) (Interface, error) {
	records, err := Read(r, 1, f)
	return records[0], err
}

// ReadMulti reads all records from a FASTA file using the generator f to validate the sequences
// Only records up to the first error are returned (along with the error)
func ReadMulti(r io.Reader, f sequence.Generator) ([]Interface, error) {
	return Read(r, 0, f)
}
