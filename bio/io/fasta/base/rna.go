package base

import (
	"io"

	"github.com/sembio/go/bio/io/fasta"
	"github.com/sembio/go/bio/sequence/immutable"
)

var _ fasta.Interface = new(Rna)

// Rna is a Fasta containing a Rna sequence
type Rna struct {
	s *Struct
}

// Header is the header line
func (x *Rna) Header() string {
	return x.s.Header()
}

// Sequence is the body lines with newlines removed
func (x *Rna) Sequence() string {
	return x.s.Sequence()
}

// ReadRna reads in a FASTA file that should contain only valid Rna letters
func ReadRna(r io.ReadCloser) (Rna, error) {
	entry, err := fasta.ReadSingle(r, func(head, body string) (fasta.Interface, error) {
		seq, err := immutable.NewRna(body)
		return New(head, seq), err
	})
	return Rna{entry.(*Struct)}, err
}

// ReadMultiRna reads in a multi-record FASTA file that should contain only valid Rna letters
func ReadMultiRna(r io.ReadCloser) ([]Rna, error) {
	entries, err := fasta.ReadMulti(r, func(head, body string) (fasta.Interface, error) {
		seq, err := immutable.NewRna(body)
		return New(head, seq), err
	})
	records := make([]Rna, len(entries))
	for i, entry := range entries {
		records[i] = Rna{entry.(*Struct)}
	}
	return records, err
}
