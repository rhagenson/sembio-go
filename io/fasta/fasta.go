package fasta

import (
	"bitbucket.org/rhagenson/bio/sequence"
)

// Fasta is the generalization of two-line FASTA format
type Fasta struct {
	header string
	body   sequence.Interface
}

// Header is the header line
func (f *Fasta) Header() string {
	return f.header
}

// Body is the body lines with newlines removed
func (f *Fasta) Body() string {
	seq, _ := f.body.Range(0, f.body.Length())
	return seq
}
