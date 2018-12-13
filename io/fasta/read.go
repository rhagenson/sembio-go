package fasta

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"bitbucket.org/rhagenson/bio/sequence"
)

// Read parses a FASTA file at r, using the generator f to validate the body
func Read(r io.Reader, f sequence.Generator) (Interface, error) {
	br := bufio.NewScanner(r)
	br.Split(bufio.ScanLines)

	header := ""
	body := ""
	for br.Scan() {
		if strings.HasPrefix(br.Text(), string(FastaHeaderPrefix)) {
			if header != "" {
				return nil, fmt.Errorf("second header line found, only one expected")
			}
			header = strings.TrimSpace(
				strings.TrimLeft(
					br.Text(),
					string(FastaHeaderPrefix),
				),
			)
		} else {
			body = body + br.Text()
		}
	}

	seq, err := f(body)

	return &Struct{
		header: header,
		body:   seq,
	}, err
}
