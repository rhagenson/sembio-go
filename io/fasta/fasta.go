package fasta

import (
	"bitbucket.org/rhagenson/bio/sequence"
)

var _ Interface = new(Struct)

// Struct is the generalization of two-line FASTA format
type Struct struct {
	header string
	seq    sequence.Interface
}

// Header is the header line
func (f *Struct) Header() string {
	return f.header
}

// Sequence is the body lines with newlines removed
func (f *Struct) Sequence() string {
	seq, _ := f.seq.Range(0, f.seq.Length())
	return seq
}
