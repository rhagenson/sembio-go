package fasta

import (
	"io"

	"github.com/bio-ext/bio-go/bio/sequence"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
)

var _ Interface = new(Dna)

// Dna is a Fasta containing a DNA sequence
type Dna struct {
	s *Struct
}

// Header is the header line
func (x *Dna) Header() string {
	return x.s.Header()
}

// Sequence is the body lines with newlines removed
func (x *Dna) Sequence() string {
	return x.s.Sequence()
}

// ReadDna reads in a FASTA file that should contain only valid Dna letters
func ReadDna(r io.ReadCloser) (Dna, error) {
	entry, err := ReadSingle(r, func(s string) (sequence.Interface, error) {
		return immutable.NewDna(s)
	})
	return Dna{entry.(*Struct)}, err
}

// ReadMultiDna reads in a multi-record FASTA file that should contain only valid Dna letters
func ReadMultiDna(r io.ReadCloser) ([]Dna, error) {
	entries, err := ReadMulti(r, func(s string) (sequence.Interface, error) {
		return immutable.NewDna(s)
	})
	records := make([]Dna, len(entries))
	for i, entry := range entries {
		records[i] = Dna{entry.(*Struct)}
	}
	return records, err
}
