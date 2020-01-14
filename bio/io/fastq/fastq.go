package fastq

import (
	"github.com/sembio/go/bio/io/fasta/base"
	"github.com/sembio/go/bio/sequence"
)

var _ Interface = new(Struct)

// Struct is the generalization of four-line FASTQ format
type Struct struct {
	f       *base.Struct
	quality string
}

// New is an Struct generator
func New(header, quality string, seq sequence.Interface) *Struct {
	return &Struct{
		f:       base.New(header, seq),
		quality: quality,
	}
}

// Header is the header line
func (x *Struct) Header() string {
	return x.f.Header()
}

// Sequence is the body lines with newlines removed
func (x *Struct) Sequence() string {
	return x.f.Sequence()
}

// Quality is the quality line
func (x *Struct) Quality() string {
	return x.quality
}
