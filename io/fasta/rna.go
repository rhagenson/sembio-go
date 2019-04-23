package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/immutable"
)

var _ Interface = new(Rna)

// Rna is a Fasta containing a Rna sequence
type Rna struct {
	*Struct
}

// ReadRna reads in a FASTA file that should contain only valid Rna letters
func ReadRna(r io.ReadCloser) (Rna, error) {
	entry, err := ReadSingle(r, func(s string) (sequence.Interface, error) {
		return immutable.NewRna(s)
	})
	return Rna{entry.(*Struct)}, err
}

// ReadMultiRna reads in a multi-record FASTA file that should contain only valid Rna letters
func ReadMultiRna(r io.ReadCloser) ([]Rna, error) {
	entries, err := ReadMulti(r, func(s string) (sequence.Interface, error) {
		return immutable.NewRna(s)
	})
	records := make([]Rna, len(entries))
	for i, entry := range entries {
		records[i] = Rna{entry.(*Struct)}
	}
	return records, err
}
