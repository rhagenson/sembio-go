package base

import (
	"io"

	"github.com/bio-ext/bio-go/bio/io/fasta"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
)

var _ fasta.Interface = new(DnaIupac)

// DnaIupac is a Fasta containing a DnaIupac sequence
type DnaIupac struct {
	s *Struct
}

// Header is the header line
func (x *DnaIupac) Header() string {
	return x.s.Header()
}

// Sequence is the body lines with newlines removed
func (x *DnaIupac) Sequence() string {
	return x.s.Sequence()
}

// ReadDnaIupac reads in a FASTA file that should contain only valid DnaIupac letters
func ReadDnaIupac(r io.ReadCloser) (DnaIupac, error) {
	entry, err := fasta.ReadSingle(r, func(head, body string) (fasta.Interface, error) {
		seq, err := immutable.NewDnaIupac(body)
		return New(head, seq), err
	})
	return DnaIupac{entry.(*Struct)}, err
}

// ReadMultiDnaIupac reads in a multi-record FASTA file that should contain only valid DnaIupac letters
func ReadMultiDnaIupac(r io.ReadCloser) ([]DnaIupac, error) {
	entries, err := fasta.ReadMulti(r, func(head, body string) (fasta.Interface, error) {
		seq, err := immutable.NewDnaIupac(body)
		return New(head, seq), err
	})
	records := make([]DnaIupac, len(entries))
	for i, entry := range entries {
		records[i] = DnaIupac{entry.(*Struct)}
	}
	return records, err
}
