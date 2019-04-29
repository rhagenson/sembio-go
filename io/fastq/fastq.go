package fastq

import (
	"bitbucket.org/rhagenson/bio/sequence"
)

var _ Interface = new(Struct)

// Struct is the generalization of four-line FASTQ format
type Struct struct {
	header  string
	seq     sequence.Interface
	quality string
}

// New is an Struct generator
func New(header, quality string, seq sequence.Interface) *Struct {
	return &Struct{
		header:  header,
		seq:     seq,
		quality: quality,
	}
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

// Quality is the quality line
func (f *Struct) Quality() string {
	return f.quality
}
