package fastq

import (
	"io"

	"github.com/bio-ext/bio-go/bio/sequence"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
)

var _ Interface = new(ProteinGapped)

// ProteinGapped is a FASTQ containing a ProteinGapped sequence
type ProteinGapped struct {
	s *Struct
}

// ReadProteinGapped reads in a FASTQ file that should contain only valid ProteinGapped letters
func ReadProteinGapped(r io.ReadCloser) (ProteinGapped, error) {
	entry, err := ReadSingle(r, func(s string) (sequence.Interface, error) {
		return immutable.NewProteinGapped(s)
	})
	return ProteinGapped{entry.(*Struct)}, err
}

// ReadMultiProteinGapped reads in a multi-record FASTQ file that should contain only valid ProteinGapped letters
func ReadMultiProteinGapped(r io.ReadCloser) ([]ProteinGapped, error) {
	entries, err := ReadMulti(r, func(s string) (sequence.Interface, error) {
		return immutable.NewProteinGapped(s)
	})
	records := make([]ProteinGapped, len(entries))
	for i, entry := range entries {
		records[i] = ProteinGapped{entry.(*Struct)}
	}
	return records, err
}

// Header is the header line
func (x *ProteinGapped) Header() string {
	return x.s.Header()
}

// Sequence is the body lines with newlines removed
func (x *ProteinGapped) Sequence() string {
	return x.s.Sequence()
}

// Quality is the quality line
func (x *ProteinGapped) Quality() string {
	return x.s.Quality()
}
