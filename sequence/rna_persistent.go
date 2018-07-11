package sequence

import (
	"errors"
	"strings"

	"bitbucket.org/rhagenson/bigr/alphabet"
	"bitbucket.org/rhagenson/bigr/helpers"
)

// RnaPersistent is the simplest, string-backed representation of RNA with
// full persistent
type RnaPersistent struct {
	seq  string
	errs []error
}

// Alphabet is the backing valid StrictRNA alphabet
func (s *RnaPersistent) Alphabet() *alphabet.RnaStrict {
	return new(alphabet.RnaStrict)
}

// Length is the number of nucleotides in the sequence
func (s *RnaPersistent) Length() uint {
	return uint(len(s.seq))
}

// Position is the nucleotide found at position n
func (s *RnaPersistent) Position(n uint) string {
	return string(s.seq[n])
}

// Range is the nucleotides found in the half-open range
func (s *RnaPersistent) Range(start, stop uint) string {
	if stop == s.Length() {
		return s.seq[start:]
	}
	return s.seq[start:stop]
}

// WithPosition mutates a sequence position
func (s *RnaPersistent) WithPosition(n uint, pos string) *RnaPersistent {
	seq := NewRnaPersistent(s.seq[:n] + pos + s.seq[n+1:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// WithRange mutates a range of sequence positions
func (s *RnaPersistent) WithRange(start, stop uint, pos string) *RnaPersistent {
	seq := NewRnaPersistent(s.seq[:start] + pos + s.seq[stop:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// NewRnaPersistent creates a new RnaPersistent instance
func NewRnaPersistent(s string) *RnaPersistent {
	seq := new(RnaPersistent)
	seq.seq = s
	seq.errs = make([]error, 1)

	acc := 0
	for _, r := range alphabet.RnaStrictLetters {
		acc += strings.Count(s, string(r))
	}
	if acc != len(s) {
		seq.errs = append(
			seq.errs,
			errors.New("string contains non-valid character"),
		)
	}

	return seq
}

// Errors returns any accumulated errors
func (s *RnaPersistent) Errors() []error {
	return s.Errors()
}

// Complement returns the base pair complement
func (s *RnaPersistent) Complement() *RnaPersistent {
	t := make([]byte, s.Length())
	for i := 0; i < len(t); i++ {
		t[i] = helpers.CompAUGC(byte(s.seq[i]))
	}
	seq := NewRnaPersistent(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// Reverse reverses the sequence
func (s *RnaPersistent) Reverse() *RnaPersistent {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i], t[l-1-i] = s.seq[l-1-i], s.seq[i]
	}
	seq := NewRnaPersistent(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// RevComp reverses and complements the sequence directly
// rather than chain the Reverse().Comp() operations together
func (s *RnaPersistent) RevComp() *RnaPersistent {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i] = helpers.CompAUGC(s.seq[l-1-i])
		t[l-1-i] = helpers.CompAUGC(s.seq[i])
	}
	seq := NewRnaPersistent(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}
