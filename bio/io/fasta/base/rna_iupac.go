package base

import (
	"io"

	"github.com/bio-ext/bio-go/bio/io/fasta"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
)

var _ fasta.Interface = new(RnaIupac)

// RnaIupac is a Fasta containing a RnaIupac sequence
type RnaIupac struct {
	s *Struct
}

// Header is the header line
func (x *RnaIupac) Header() string {
	return x.s.Header()
}

// Sequence is the body lines with newlines removed
func (x *RnaIupac) Sequence() string {
	return x.s.Sequence()
}

// ReadRnaIupac reads in a FASTA file that should contain only valid RnaIupac letters
func ReadRnaIupac(r io.ReadCloser) (RnaIupac, error) {
	entry, err := fasta.ReadSingle(r, func(head, body string) (fasta.Interface, error) {
		seq, err := immutable.NewRnaIupac(body)
		return New(head, seq), err
	})
	return RnaIupac{entry.(*Struct)}, err
}

// ReadMultiRnaIupac reads in a multi-record FASTA file that should contain only valid RnaIupac letters
func ReadMultiRnaIupac(r io.ReadCloser) ([]RnaIupac, error) {
	entries, err := fasta.ReadMulti(r, func(head, body string) (fasta.Interface, error) {
		seq, err := immutable.NewRnaIupac(body)
		return New(head, seq), err
	})
	records := make([]RnaIupac, len(entries))
	for i, entry := range entries {
		records[i] = RnaIupac{entry.(*Struct)}
	}
	return records, err
}
