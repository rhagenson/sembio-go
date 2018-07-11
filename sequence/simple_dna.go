package sequence

import (
	"errors"
	"strings"

	"bitbucket.org/rhagenson/bigr/alphabet"
	"bitbucket.org/rhagenson/bigr/helpers"
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
	seq := NewSimpleDna(s.seq[:n] + pos + s.seq[n+1:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// WithRange mutates a range of sequence positions
func (s *SimpleDna) WithRange(start, stop uint, pos string) *SimpleDna {
	seq := NewSimpleDna(s.seq[:start] + pos + s.seq[stop:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// NewSimpleDna creates a new SimpleDna instance if the input is valid DNA
func NewSimpleDna(s string) *SimpleDna {
	seq := new(SimpleDna)
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
func (s *SimpleDna) Errors() []error {
	return s.Errors()
}

// Complement returns the base pair complement
func (s *SimpleDna) Complement() *SimpleDna {
	t := make([]byte, s.Length())
	for i := 0; i < len(t); i++ {
		t[i] = helpers.CompATGC(byte(s.seq[i]))
	}
	seq := NewSimpleDna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// Reverse reverses the sequence
func (s *SimpleDna) Reverse() *SimpleDna {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i], t[l-1-i] = s.seq[l-1-i], s.seq[i]
	}
	seq := NewSimpleDna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// RevComp reverses and complements the sequence directly
// rather than chain the Reverse().Comp() operations together
func (s *SimpleDna) RevComp() *SimpleDna {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i] = helpers.CompATGC(s.seq[l-1-i])
		t[l-1-i] = helpers.CompATGC(s.seq[i])
	}
	seq := NewSimpleDna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}
