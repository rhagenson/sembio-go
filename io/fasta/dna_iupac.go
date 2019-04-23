package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/immutable"
)

var _ Interface = new(DnaIupac)

// DnaIupac is a Fasta containing a DnaIupac sequence
type DnaIupac struct {
	*Struct
}

// ReadDnaIupac reads in a FASTA file that should contain only valid DnaIupac letters
func ReadDnaIupac(r io.ReadCloser) (DnaIupac, error) {
	entry, err := ReadSingle(r, func(s string) (sequence.Interface, error) {
		return immutable.NewDnaIupac(s)
	})
	return DnaIupac{entry.(*Struct)}, err
}

// ReadMultiDnaIupac reads in a multi-record FASTA file that should contain only valid DnaIupac letters
func ReadMultiDnaIupac(r io.ReadCloser) ([]DnaIupac, error) {
	entries, err := ReadMulti(r, func(s string) (sequence.Interface, error) {
		return immutable.NewDnaIupac(s)
	})
	records := make([]DnaIupac, len(entries))
	for i, entry := range entries {
		records[i] = DnaIupac{entry.(*Struct)}
	}
	return records, err
}
