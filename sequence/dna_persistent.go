package sequence

import (
	"errors"
	"strings"

	"bitbucket.org/rhagenson/bigr/alphabet"
	"bitbucket.org/rhagenson/bigr/helpers"
)

// DnaPersistent is the simplest, string-backed representation of DNA with
// full persistent
type DnaPersistent struct {
	seq  string
	errs []error
}

// Alphabet is the backing valid StrictDNA alphabet
func (s *DnaPersistent) Alphabet() *alphabet.DnaStrict {
	return new(alphabet.DnaStrict)
}

// Length is the number of nucleotides in the sequence
func (s *DnaPersistent) Length() uint {
	return uint(len(s.seq))
}

// Position is the nucleotide found at position n
func (s *DnaPersistent) Position(n uint) string {
	return string(s.seq[n])
}

// Range is the nucleotides found in the half-open range
func (s *DnaPersistent) Range(start, stop uint) string {
	if stop == s.Length() {
		return s.seq[start:]
	}
	return s.seq[start:stop]
}

// WithPosition mutates a sequence position
func (s *DnaPersistent) WithPosition(n uint, pos string) *DnaPersistent {
	seq := NewDnaPersistent(s.seq[:n] + pos + s.seq[n+1:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// WithRange mutates a range of sequence positions
func (s *DnaPersistent) WithRange(start, stop uint, pos string) *DnaPersistent {
	seq := NewDnaPersistent(s.seq[:start] + pos + s.seq[stop:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// NewDnaPersistent creates a new DnaPersistent instance if the input is valid DNA
func NewDnaPersistent(s string) *DnaPersistent {
	seq := new(DnaPersistent)
	seq.seq = s
	seq.errs = make([]error, 1)

	acc := 0
	for _, r := range alphabet.DnaStrictLetters {
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
func (s *DnaPersistent) Errors() []error {
	return s.Errors()
}

// Complement returns the base pair complement
func (s *DnaPersistent) Complement() *DnaPersistent {
	t := make([]byte, s.Length())
	for i := 0; i < len(t); i++ {
		t[i] = helpers.CompATGC(byte(s.seq[i]))
	}
	seq := NewDnaPersistent(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// Reverse reverses the sequence
func (s *DnaPersistent) Reverse() *DnaPersistent {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i], t[l-1-i] = s.seq[l-1-i], s.seq[i]
	}
	seq := NewDnaPersistent(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// RevComp reverses and complements the sequence directly
// rather than chain the Reverse().Comp() operations together
func (s *DnaPersistent) RevComp() *DnaPersistent {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i] = helpers.CompATGC(s.seq[l-1-i])
		t[l-1-i] = helpers.CompATGC(s.seq[i])
	}
	seq := NewDnaPersistent(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}
