package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
)

var _ Interface = new(Dna)

// Dna is a Fasta containing a DNA sequence
type Dna struct {
	*Struct
}

// ReadDna reads in a FASTA file that should contain only valid DNA letters
func ReadDna(r io.Reader) (Dna, error) {
	s, err := Read(r, func(s string) (sequence.Interface, error) {
		return sequence.NewDna(s)
	})
	return Dna{s.(*Struct)}, err
}
