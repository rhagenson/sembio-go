package fasta

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"bitbucket.org/rhagenson/bio/sequence"
)

// ReadSingle parses a single-record FASTA file using the generator f to validate the sequence
func ReadSingle(r io.Reader, f sequence.Generator) (Interface, error) {
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
		seq:    seq,
	}, err
}

// ReadMulti parses a multi-record FASTA file using the generator f to validate the sequences
func ReadMulti(r io.Reader, f sequence.Generator) ([]Interface, error) {
	seqs := make([]Interface, 1)
	br := bufio.NewScanner(r)
	br.Split(bufio.ScanLines)

	header := ""
	body := ""
	for br.Scan() {
		if strings.HasPrefix(br.Text(), string(FastaHeaderPrefix)) {
			if header != "" {
				seq, err := f(body)
				if err != nil {
					return seqs, err
				}
				seqs = append(seqs, &Struct{
					header: header,
					seq:    seq,
				})
			}
			header = strings.TrimSpace(
				strings.TrimLeft(
					br.Text(),
					string(FastaHeaderPrefix),
				),
			)
			body = ""
		} else {
			body = body + br.Text()
		}
	}
	return seqs, nil
}
