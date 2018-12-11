package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/persistent"
)

var _ Interface = new(ProteinGapped)

// ProteinGapped is a Fasta containing a ProteinGapped sequence
type ProteinGapped struct {
	*Struct
}

// ReadProteinGapped reads in a FASTA file that should contain only valid ProteinGapped letters
func ReadProteinGapped(r io.Reader) (ProteinGapped, error) {
	s, err := Read(r, func(s string) (sequence.Interface, error) {
		return persistent.NewProteinGapped(s)
	})
	return ProteinGapped{s.(*Struct)}, err
}
