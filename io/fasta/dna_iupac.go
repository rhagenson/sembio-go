package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
)

var _ Interface = new(DnaIupac)

// DnaIupac is a Fasta containing a DnaIupac sequence
type DnaIupac struct {
	*Struct
}

// ReadDnaIupac reads in a FASTA file that should contain only valid DnaIupac letters
func ReadDnaIupac(r io.Reader) (DnaIupac, error) {
	s, err := Read(r, func(s string) (sequence.Interface, error) {
		return sequence.NewDnaIupac(s)
	})
	return DnaIupac{s.(*Struct)}, err
}