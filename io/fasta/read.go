package fasta

import (
	"bufio"
	"io"
	"strings"

	"bitbucket.org/rhagenson/bigr/sequence"
)

// Read reads in a FASTA file at r which should use the alphabet of t
func Read(r io.Reader, f func(string) (sequence.Interface, error)) (Interface, error) {
	br := bufio.NewScanner(r)
	br.Split(bufio.ScanLines)

	header := ""
	body := ""
	for br.Scan() {
		if strings.HasPrefix(br.Text(), string(fastaHeaderPrefix)) {
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
	}, err
}
