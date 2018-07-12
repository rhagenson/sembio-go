package persistent

import (
	"errors"
	"strings"

	"bitbucket.org/rhagenson/bigr/alphabet"
	"bitbucket.org/rhagenson/bigr/helpers"
)

// RnaIupac is the simplest, string-backed representation of RNA with
// full persistence
type RnaIupac struct {
	seq  string
	errs []error
}

// Alphabet is the backing valid StrictRNA alphabet
func (s *RnaIupac) Alphabet() *alphabet.RnaIupac {
	return new(alphabet.RnaIupac)
}

// Length is the number of nucleotides in the sequence
func (s *RnaIupac) Length() uint {
	return uint(len(s.seq))
}

// Position is the nucleotide found at position n
func (s *RnaIupac) Position(n uint) string {
	return string(s.seq[n])
}

// Range is the nucleotides found in the half-open range
func (s *RnaIupac) Range(start, stop uint) string {
	if stop == s.Length() {
		return s.seq[start:]
	}
	return s.seq[start:stop]
}

// WithPosition mutates a sequence position
func (s *RnaIupac) WithPosition(n uint, pos string) *RnaIupac {
	seq := NewRnaIupac(s.seq[:n] + pos + s.seq[n+1:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// WithRange mutates a range of sequence positions
func (s *RnaIupac) WithRange(start, stop uint, pos string) *RnaIupac {
	seq := NewRnaIupac(s.seq[:start] + pos + s.seq[stop:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// NewRnaIupac creates a new RnaIupac instance
func NewRnaIupac(s string) *RnaIupac {
	seq := new(RnaIupac)
	seq.seq = s
	seq.errs = make([]error, 0)

	acc := 0
	for _, r := range alphabet.RnaIupacLetters {
		acc += strings.Count(s, string(r))
	}
	if acc != len(s) {
		seq.errs = append(
			seq.errs,
			errors.New("sequence contains invalid character(s) on creation"),
		)
	}

	return seq
}

// Errors returns any accumulated errors
func (s *RnaIupac) Errors() []error {
	return s.Errors()
}

// Complement returns the base pair complement
func (s *RnaIupac) Complement() *RnaIupac {
	t := make([]byte, s.Length())
	for i := 0; i < len(t); i++ {
		t[i] = helpers.CompAUGC(byte(s.seq[i]))
	}
	seq := NewRnaIupac(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// Reverse reverses the sequence
func (s *RnaIupac) Reverse() *RnaIupac {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i], t[l-1-i] = s.seq[l-1-i], s.seq[i]
	}
	seq := NewRnaIupac(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// RevComp reverses and complements the sequence directly
// rather than chain the Reverse().Comp() operations together
func (s *RnaIupac) RevComp() *RnaIupac {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i] = helpers.CompAUGC(s.seq[l-1-i])
		t[l-1-i] = helpers.CompAUGC(s.seq[i])
	}
	seq := NewRnaIupac(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}
