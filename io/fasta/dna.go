package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/immutable"
)

var _ Interface = new(Dna)

// Dna is a Fasta containing a DNA sequence
type Dna struct {
	*Struct
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
