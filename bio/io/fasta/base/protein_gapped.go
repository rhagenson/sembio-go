package base

import (
	"io"

	"github.com/bio-ext/bio-go/bio/io/fasta"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
)

var _ fasta.Interface = new(ProteinGapped)

// ProteinGapped is a Fasta containing a ProteinGapped sequence
type ProteinGapped struct {
	s *Struct
}

// Header is the header line
func (x *ProteinGapped) Header() string {
	return x.s.Header()
}

// Sequence is the body lines with newlines removed
func (x *ProteinGapped) Sequence() string {
	return x.s.Sequence()
}

// ReadProteinGapped reads in a FASTA file that should contain only valid ProteinGapped letters
func ReadProteinGapped(r io.ReadCloser) (ProteinGapped, error) {
	entry, err := fasta.ReadSingle(r, func(head, body string) (fasta.Interface, error) {
		seq, err := immutable.NewProteinGapped(body)
		return New(head, seq), err
	})
	return ProteinGapped{entry.(*Struct)}, err
}

// ReadMultiProteinGapped reads in a multi-record FASTA file that should contain only valid ProteinGapped letters
func ReadMultiProteinGapped(r io.ReadCloser) ([]ProteinGapped, error) {
	entries, err := fasta.ReadMulti(r, func(head, body string) (fasta.Interface, error) {
		seq, err := immutable.NewProteinGapped(body)
		return New(head, seq), err
	})
	records := make([]ProteinGapped, len(entries))
	for i, entry := range entries {
		records[i] = ProteinGapped{entry.(*Struct)}
	}
	return records, err
}
