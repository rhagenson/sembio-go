package fasta

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"bitbucket.org/rhagenson/bigr/sequence"
)

// Read parses a FASTA file at r, using the genrator f to validate the body
func Read(r io.Reader, f sequence.Generator) (Interface, error) {
	br := bufio.NewScanner(r)
	br.Split(bufio.ScanLines)

	header := ""
	body := ""
	for br.Scan() {
		if strings.HasPrefix(br.Text(), string(fastaHeaderPrefix)) {
			if header != "" {
				return nil, fmt.Errorf("fasta should only have one header line")
			}
			header = header + strings.TrimSpace(
				strings.TrimLeft(
					br.Text(),
					string(fastaHeaderPrefix),
				),
			)
		} else {
			body = body + br.Text()
		}
	}

	seq, err := f(body)

	return &Fasta{
		header: header,
		body:   seq,
	}, fmt.Errorf("invalid sequence: %v", err)
}
