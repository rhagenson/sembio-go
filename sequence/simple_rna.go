package sequence

import (
	"errors"
	"strings"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

// SimpleRna is the simplest, string-backed representation of DNA
type SimpleRna struct {
	seq  string
	errs []error
}

// Alphabet is the backing valid StrictRNA alphabet
func (s *SimpleRna) Alphabet() *alphabet.RNAStrict {
	return new(alphabet.RNAStrict)
}

// Length is the number of nucleotides in the sequence
func (s *SimpleRna) Length() uint {
	return uint(len(s.seq))
}

// Position is the nucleotide found at position n
func (s *SimpleRna) Position(n uint) string {
	return string(s.seq[n])
}

// Range is the nucleotides found in the half-open range
func (s *SimpleRna) Range(start, stop uint) string {
	if stop == s.Length() {
		return s.seq[start:]
	}
	return s.seq[start:stop]
}

// WithPosition mutates a sequence position
func (s *SimpleRna) WithPosition(n uint, pos string) *SimpleRna {
	seq := NewSimpleRna(s.seq[:n] + pos + s.seq[n+1:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// WithRange mutates a range of sequence positions
func (s *SimpleRna) WithRange(start, stop uint, pos string) *SimpleRna {
	seq := NewSimpleRna(s.seq[:start] + pos + s.seq[stop:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// NewSimpleRna creates a new SimpleRna instance
func NewSimpleRna(s string) *SimpleRna {
	seq := new(SimpleRna)
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
func (s *SimpleRna) Errors() []error {
	return s.Errors()
}
