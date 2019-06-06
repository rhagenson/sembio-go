package fasta

import (
	"io"

	"github.com/rhagenson/bio-go/bio/sequence"
	"github.com/rhagenson/bio-go/bio/sequence/immutable"
)

var _ Interface = new(ProteinGapped)

// ProteinGapped is a Fasta containing a ProteinGapped sequence
type ProteinGapped struct {
	*Struct
}

// ReadProteinGapped reads in a FASTA file that should contain only valid ProteinGapped letters
func ReadProteinGapped(r io.ReadCloser) (ProteinGapped, error) {
	entry, err := ReadSingle(r, func(s string) (sequence.Interface, error) {
		return immutable.NewProteinGapped(s)
	})
	return ProteinGapped{entry.(*Struct)}, err
}

// ReadMultiProteinGapped reads in a multi-record FASTA file that should contain only valid ProteinGapped letters
func ReadMultiProteinGapped(r io.ReadCloser) ([]ProteinGapped, error) {
	entries, err := ReadMulti(r, func(s string) (sequence.Interface, error) {
		return immutable.NewProteinGapped(s)
	})
	records := make([]ProteinGapped, len(entries))
	for i, entry := range entries {
		records[i] = ProteinGapped{entry.(*Struct)}
	}
	return records, err
}
