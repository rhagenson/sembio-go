package sequence

import (
	"errors"
	"strings"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

// SimpleDna is the simplest, string-backed representation of DNA
type SimpleDna struct {
	seq  string
	errs []error
}

// Alphabet is the backing valid StrictDNA alphabet
func (s *SimpleDna) Alphabet() *alphabet.DNAStrict {
	return new(alphabet.DNAStrict)
}

// Length is the number of nucleotides in the sequence
func (s *SimpleDna) Length() uint {
	return uint(len(s.seq))
}

// Position is the nucleotide found at position n
func (s *SimpleDna) Position(n uint) string {
	return string(s.seq[n])
}

// Range is the nucleotides found in the half-open range
func (s *SimpleDna) Range(start, stop uint) string {
	if stop == s.Length() {
		return s.seq[start:]
	}
	return s.seq[start:stop]
}

// WithPosition mutates a sequence position
func (s *SimpleDna) WithPosition(n uint, pos string) *SimpleDna {
	return NewSimpleDna(s.seq[:n] + pos + s.seq[n+1:])
}

// WithRange mutates a range of sequence positions
func (s *SimpleDna) WithRange(start, stop uint, pos string) *SimpleDna {
	return NewSimpleDna(s.seq[:start] + pos + s.seq[stop:])
}

// NewSimpleDna creates a new SimpleDna instance if the input is valid DNA
func NewSimpleDna(s string) (seq *SimpleDna) {
	acc := 0
	for _, r := range alphabet.DnaStrictLetters {
		acc += strings.Count(s, string(r))
	}
	if acc != len(s) {
		seq.errs = append(seq.errs, errors.New("string contains non-valid character"))
	} else {
		seq.seq = s
	}
	return
}

// Errors returns any accumulated errors
func (s *SimpleDna) Errors() []error {
	return s.Errors()
}
