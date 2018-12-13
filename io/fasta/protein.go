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
func ReadProtein(r io.Reader) (Protein, error) {
	s, err := Read(r, func(s string) (sequence.Interface, error) {
		return immutable.NewProtein(s)
	})
	return Protein{s.(*Struct)}, err
}
