package fasta

import (
	"io"

	"github.com/bio-ext/bio-go/bio/sequence"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
)

var _ Interface = new(Protein)

// Protein is a Fasta containing a Protein sequence
type Protein struct {
	s *Struct
}

// Header is the header line
func (x *Protein) Header() string {
	return x.s.Header()
}

// Sequence is the body lines with newlines removed
func (x *Protein) Sequence() string {
	return x.s.Sequence()
}

// ReadProtein reads in a FASTA file that should contain only valid Protein letters
func ReadProtein(r io.ReadCloser) (Protein, error) {
	entry, err := ReadSingle(r, func(s string) (sequence.Interface, error) {
		return immutable.NewProtein(s)
	})
	return Protein{entry.(*Struct)}, err
}

// ReadMultiProtein reads in a multi-record FASTA file that should contain only valid Protein letters
func ReadMultiProtein(r io.ReadCloser) ([]Protein, error) {
	entries, err := ReadMulti(r, func(s string) (sequence.Interface, error) {
		return immutable.NewProtein(s)
	})
	records := make([]Protein, len(entries))
	for i, entry := range entries {
		records[i] = Protein{entry.(*Struct)}
	}
	return records, err
}
