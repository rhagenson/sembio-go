package fasta

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"bitbucket.org/rhagenson/bio/sequence"
	"github.com/pkg/errors"
)

// Read will read FASTA and produce n records.
// If n == 0, all records are read
func Read(r io.Reader, f sequence.Generator, n uint) (<-chan Interface, <-chan error) {
	ch1 := make(chan Interface, n)
	ch2 := make(chan error, n)
	count := uint(0)

	go func() {
		br := bufio.NewScanner(r)
		br.Split(bufio.ScanLines)
		header := ""
		body := ""
		for br.Scan() {
			if strings.HasPrefix(br.Text(), string(FastaHeaderPrefix)) {
				if body != "" && header != "" {
					seq, err := f(body)
					record := New(header, seq)
					fmt.Printf("Ryan: %q\n", record.Header())
					ch1 <- record
					ch2 <- errors.Wrap(err, fmt.Sprintf("error in record %q", header))
					count++
					if n != 0 && count == n {
						goto end
					}
					header = ""
					body = ""
				} else {
					header = strings.TrimSpace(
						strings.TrimLeft(
							br.Text(),
							string(FastaHeaderPrefix),
						),
					)
				}
			} else {
				body = body + br.Text()
			}
		}
	end:
		close(ch1)
		close(ch2)
		return
	}()
	return ch1, ch2
}

// ReadSingle reads a single records from a FASTA file using the generator f to validate the sequence
func ReadSingle(r io.Reader, f sequence.Generator) (Interface, error) {
	records, errs := Read(r, f, 1)
	return <-records, <-errs
}

// ReadMulti reads all records from a FASTA file using the generator f to validate the sequences
// Only records up to the first error are returned (along with the error)
func ReadMulti(r io.Reader, f sequence.Generator) ([]Interface, error) {
	records, errs := Read(r, f, 0)
	bufRec := make([]Interface, 1)
	select {
	case record := <-records:
		bufRec = append(bufRec, record)
	case err := <-errs:
		if err != nil {
			return bufRec, err
		}
	}
	return bufRec, nil
}
