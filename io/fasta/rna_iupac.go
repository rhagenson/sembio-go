package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
)

var _ Interface = new(RnaIupac)

// RnaIupac is a Fasta containing a RnaIupac sequence
type RnaIupac struct {
	*Struct
}

// ReadRnaIupac reads in a FASTA file that should contain only valid RnaIupac letters
func ReadRnaIupac(r io.Reader) (RnaIupac, error) {
	s, err := Read(r, func(s string) (sequence.Interface, error) {
		return sequence.NewRnaIupac(s)
	})
	return RnaIupac{s.(*Struct)}, err
}