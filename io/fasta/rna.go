package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
)

var _ Interface = new(Rna)

// Rna is a Fasta containing a Rna sequence
type Rna struct {
	*Struct
}

// ReadRna reads in a FASTA file that should contain only valid Rna letters
func ReadRna(r io.Reader) (Rna, error) {
	s, err := Read(r, func(s string) (sequence.Interface, error) {
		return sequence.NewRna(s)
	})
	return Rna{s.(*Struct)}, err
}
