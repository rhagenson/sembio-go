package persistent

import (
	"errors"
	"fmt"
	"strings"

	"bitbucket.org/rhagenson/bigr/alphabet/simple"
	"bitbucket.org/rhagenson/bigr/helpers/complement"
)

// Rna is the simplest, string-backed representation of RNA with
// full persistence
type Rna struct {
	seq  string
	errs []error
}

// Alphabet is the backing valid StrictRNA alphabet
func (s *Rna) Alphabet() *simple.Rna {
	return new(simple.Rna)
}

// Length is the number of nucleotides in the sequence
func (s *Rna) Length() uint {
	return uint(len(s.seq))
}

// Position is the nucleotide found at position n
func (s *Rna) Position(n uint) (string, error) {
	if n < uint(len(s.seq)) {
		return string(s.seq[n]), nil
	}
	s.errs = append(
		s.errs,
		fmt.Errorf("requested impossible position [%d]", n),
	)
	return "", s.errs[len(s.errs)-1]
}

// Range is the nucleotides found in the half-open range
func (s *Rna) Range(start, stop uint) (string, error) {
	if stop == s.Length() {
		return s.seq[start:], nil
	} else if start < stop && stop < s.Length() {
		return s.seq[start:stop], nil
	}

	s.errs = append(
		s.errs,
		fmt.Errorf("requested impossible range [%d:%d]", start, stop),
	)
	return "", s.errs[len(s.errs)-1]
}

// WithPosition mutates a sequence position
func (s *Rna) WithPosition(n uint, pos string) (*Rna, error) {
	seq, err := NewRna(s.seq[:n] + pos + s.seq[n+1:])
	seq.errs = append(s.errs, seq.errs...)
	return seq, err
}

// WithRange mutates a range of sequence positions
func (s *Rna) WithRange(start, stop uint, pos string) (*Rna, error) {
	seq, err := NewRna(s.seq[:start] + pos + s.seq[stop:])
	seq.errs = append(s.errs, seq.errs...)
	return seq, err
}

// NewRna creates a new Rna instance
func NewRna(s string) (*Rna, error) {
	seq := new(Rna)
	seq.seq = s
	seq.errs = make([]error, 0)

	acc := 0
	for _, r := range simple.RnaLetters {
		acc += strings.Count(s, string(r))
	}
	if acc != len(s) {
		seq.errs = append(
			seq.errs,
			errors.New("sequence contains invalid character(s) on creation"),
		)
		return seq, seq.errs[len(seq.errs)-1]
	}
	return seq, nil
}

// Errors returns any accumulated errors
func (s *Rna) Errors() []error {
	errs := make([]error, len(s.errs))
	copy(errs, s.errs)
	return errs
}

// Complement returns the base pair complement
func (s *Rna) Complement() (*Rna, error) {
	t := make([]byte, s.Length())
	for i := 0; i < len(t); i++ {
		t[i] = complement.Augc(byte(s.seq[i]))
	}
	seq, err := NewRna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq, err
}

// Reverse reverses the sequence
func (s *Rna) Reverse() (*Rna, error) {
	l := int(s.Length())
	t := make([]byte, l)
	if st, err := s.Range(0, s.Length()); err != nil {
		t = []byte(st)
	} else {
		s.errs = append(s.errs, err)
		return s, err
	}
	for i := 0; i < l/2; i++ {
		t[i], t[l-1-i] = s.seq[l-1-i], s.seq[i]
	}
	seq, err := NewRna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq, err
}

// RevComp reverses and complements the sequence directly
// rather than chain the Reverse().Comp() operations together
func (s *Rna) RevComp() (*Rna, error) {
	l := int(s.Length())
	t := make([]byte, l)
	if st, err := s.Range(0, s.Length()); err != nil {
		t = []byte(st)
	} else {
		s.errs = append(s.errs, err)
		return s, err
	}
	for i := 0; i < l/2; i++ {
		t[i] = complement.Augc(s.seq[l-1-i])
		t[l-1-i] = complement.Augc(s.seq[i])
	}
	seq, err := NewRna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq, err
}
