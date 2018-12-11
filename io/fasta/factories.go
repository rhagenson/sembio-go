package fasta

import (
	"io"

	"bitbucket.org/rhagenson/bigr/sequence"
)

// ReadDna reads in a FASTA file that should contain only valid DNA letters
func ReadDna(r io.Reader) (Interface, error) {
	return Read(r, func(s string) (sequence.Interface, error) {
		return sequence.NewDna(s)
	})
}

// ReadDnaIupac reads in a FASTA file that should contain
// only valid IUPAC DNA letters
func ReadDnaIupac(r io.Reader) (Interface, error) {
	return Read(r, func(s string) (sequence.Interface, error) {
		return sequence.NewDnaIupac(s)
	})
}

// ReadRna reads in a FASTA file that should contain only valid DNA letters
func ReadRna(r io.Reader) (Interface, error) {
	return Read(r, func(s string) (sequence.Interface, error) {
		return sequence.NewRna(s)
	})
}

// ReadRnaIupac reads in a FASTA file that should contain
// only valid IUPAC RNA letters
func ReadRnaIupac(r io.Reader) (Interface, error) {
	return Read(r, func(s string) (sequence.Interface, error) {
		return sequence.NewRnaIupac(s)
	})
}