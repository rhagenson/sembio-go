package fasta

import (
	"bufio"
	"io"
	"strings"
)

// Read reads n records from a FASTA file using the generator satisfy Interface
// Only records up to the first error are returned (along with the error)
func Read(r io.Reader, n uint, f func(head, body string) (Interface, error)) ([]Interface, error) {
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
		case line[0] == HeaderPrefix: // Header
			if seq.Len() == 0 {
				header = line
			} else {
				record, err := f(header, seq.String())
				if err != nil {
					return records, err
				}
				records = append(records, record)
				count++
				if n != 0 && count == n {
					return records, nil
				}
				header = ""
				seq.Reset()
			}
		default: // Sequence
			seq.WriteString(line)
		}
	}
	if header != "" && seq.String() != "" {
		record, err := f(header, seq.String())
		if err != nil {
			return records, err
		}
		records = append(records, record)
	}
	return records, nil
}

// ReadSingle reads a single records from a FASTA file using the generator f to validate the sequence
func ReadSingle(r io.Reader, f func(head, body string) (Interface, error)) (Interface, error) {
	records, err := Read(r, 1, f)
	return records[0], err
}

// ReadMulti reads all records from a FASTA file using the generator f to validate the sequences
// Only records up to the first error are returned (along with the error)
func ReadMulti(r io.Reader, f func(head, body string) (Interface, error)) ([]Interface, error) {
	return Read(r, 0, f)
}
