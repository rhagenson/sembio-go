package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/immutable"
)

var _ Interface = new(Protein)

// Protein is a Fasta containing a Protein sequence
type Protein struct {
	*Struct
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
