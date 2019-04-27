package fasta

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"bitbucket.org/rhagenson/bio/sequence"
)

// ReadSingle reads a single records from a FASTA file using the generator f to validate the sequence
func ReadSingle(r io.ReadCloser, f sequence.Generator) (Interface, error) {
	br := bufio.NewScanner(r)
	br.Split(bufio.ScanLines)
	record := new(record)
	for br.Scan() {
		line := strings.TrimSpace(br.Text())
		switch {
		case line == "": // Skip blank lines
			continue
		case line[0] == FastaHeaderPrefix: // Header
			if len(record.sequence) == 0 {
				record.header = line
			} else {
				return nil, errors.New("expected one record in FASTA input, got multiple")
			}
		default: // Sequence
			record.sequence = record.sequence + line
		}
	}
	seqx, err := f(record.sequence)
	return New(record.header, seqx), err
}

// ReadMulti reads all records from a FASTA file using the generator f to validate the sequences
// Only records up to the first error are returned (along with the error)
func ReadMulti(r io.ReadCloser, f sequence.Generator) ([]Interface, error) {
	br := bufio.NewScanner(r)
	br.Split(bufio.ScanLines)
	accum := make([]record, 0)
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
				accum = append(accum, record{
					header,
					seq.String(),
				})
			}
		default: // Sequence
			seq.WriteString(line)
		}
	}
	records := make([]Interface, len(accum))
	for i, r := range accum {
		seqx, err := f(r.sequence)
		if err != nil {
			return nil, err
		}
		records[i] = New(r.header, seqx)
	}
	return records, nil
}

type record struct {
	header   string
	sequence string
}
