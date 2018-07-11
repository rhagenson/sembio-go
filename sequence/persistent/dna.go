package persistent

import (
	"errors"
	"strings"

	"bitbucket.org/rhagenson/bigr/alphabet"
	"bitbucket.org/rhagenson/bigr/helpers"
)

// Dna is the simplest, string-backed representation of DNA with
// full persistence
type Dna struct {
	seq  string
	errs []error
}

// Alphabet is the backing valid StrictDNA alphabet
func (s *Dna) Alphabet() *alphabet.DnaStrict {
	return new(alphabet.DnaStrict)
}

// Length is the number of nucleotides in the sequence
func (s *Dna) Length() uint {
	return uint(len(s.seq))
}

// Position is the nucleotide found at position n
func (s *Dna) Position(n uint) string {
	return string(s.seq[n])
}

// Range is the nucleotides found in the half-open range
func (s *Dna) Range(start, stop uint) string {
	if stop == s.Length() {
		return s.seq[start:]
	}
	return s.seq[start:stop]
}

// WithPosition mutates a sequence position
func (s *Dna) WithPosition(n uint, pos string) *Dna {
	seq := NewDna(s.seq[:n] + pos + s.seq[n+1:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// WithRange mutates a range of sequence positions
func (s *Dna) WithRange(start, stop uint, pos string) *Dna {
	seq := NewDna(s.seq[:start] + pos + s.seq[stop:])
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// NewDna creates a new Dna instance if the input is valid DNA
func NewDna(s string) *Dna {
	seq := new(Dna)
	seq.seq = s
	seq.errs = make([]error, 0)

	acc := 0
	for _, r := range alphabet.DnaStrictLetters {
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
func (s *Dna) Errors() []error {
	return s.Errors()
}

// Complement returns the base pair complement
func (s *Dna) Complement() *Dna {
	t := make([]byte, s.Length())
	for i := 0; i < len(t); i++ {
		t[i] = helpers.CompATGC(byte(s.seq[i]))
	}
	seq := NewDna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// Reverse reverses the sequence
func (s *Dna) Reverse() *Dna {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i], t[l-1-i] = s.seq[l-1-i], s.seq[i]
	}
	seq := NewDna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}

// RevComp reverses and complements the sequence directly
// rather than chain the Reverse().Comp() operations together
func (s *Dna) RevComp() *Dna {
	l := int(s.Length())
	t := []byte(s.Range(0, s.Length()))
	for i := 0; i < l/2; i++ {
		t[i] = helpers.CompATGC(s.seq[l-1-i])
		t[l-1-i] = helpers.CompATGC(s.seq[i])
	}
	seq := NewDna(string(t))
	seq.errs = append(s.errs, seq.errs...)
	return seq
}
