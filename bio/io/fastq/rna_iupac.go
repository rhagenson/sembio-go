package fastq

import (
	"io"

	"github.com/bio-ext/bio-go/bio/sequence"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
)

var _ Interface = new(RnaIupac)

// RnaIupac is a FASTQ containing a RnaIupac sequence
type RnaIupac struct {
	s *Struct
}

// ReadRnaIupac reads in a FASTQ file that should contain only valid RnaIupac letters
func ReadRnaIupac(r io.ReadCloser) (RnaIupac, error) {
	entry, err := ReadSingle(r, func(s string) (sequence.Interface, error) {
		return immutable.NewRnaIupac(s)
	})
	return RnaIupac{entry.(*Struct)}, err
}

// ReadMultiRnaIupac reads in a multi-record FASTQ file that should contain only valid RnaIupac letters
func ReadMultiRnaIupac(r io.ReadCloser) ([]RnaIupac, error) {
	entries, err := ReadMulti(r, func(s string) (sequence.Interface, error) {
		return immutable.NewRnaIupac(s)
	})
	records := make([]RnaIupac, len(entries))
	for i, entry := range entries {
		records[i] = RnaIupac{entry.(*Struct)}
	}
	return records, err
}

// Header is the header line
func (x *RnaIupac) Header() string {
	return x.s.Header()
}

// Sequence is the body lines with newlines removed
func (x *RnaIupac) Sequence() string {
	return x.s.Sequence()
}

// Quality is the quality line
func (x *RnaIupac) Quality() string {
	return x.s.Quality()
}
