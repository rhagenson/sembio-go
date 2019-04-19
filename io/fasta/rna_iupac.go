package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/immutable"
)

var _ Interface = new(RnaIupac)

// RnaIupac is a Fasta containing a RnaIupac sequence
type RnaIupac struct {
	*Struct
}

// ReadRnaIupac reads in a FASTA file that should contain only valid RnaIupac letters
func ReadRnaIupac(r io.Reader) (RnaIupac, error) {
	entry, err := ReadSingle(r, func(s string) (sequence.Interface, error) {
		return immutable.NewRnaIupac(s)
	})
	return RnaIupac{entry.(*Struct)}, err
}

// ReadMultiRnaIupac reads in a multi-record FASTA file that should contain only valid RnaIupac letters
func ReadMultiRnaIupac(r io.Reader) ([]RnaIupac, error) {
	entries, err := ReadMulti(r, func(s string) (sequence.Interface, error) {
		return immutable.NewRnaIupac(s)
	})
	records := make([]RnaIupac, len(entries))
	for i, entry := range entries {
		records[i] = RnaIupac{entry.(*Struct)}
	}
	return records, err
}
