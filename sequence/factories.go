package sequence

import (
	"bitbucket.org/rhagenson/bigr/alphabet"
)

// Generator funcs create a particular *Sequence from a given string
// The string may be invalid and therefore an error can result.
type Generator func(string) (*Sequence, error)

// NewDna generates a new sequence that validates against the Dna alphabet
func NewDna(s string) (*Sequence, error) {
	n := New(
		s,
		map[string]SeqFunc{
			"Reverse":    reverseDna,
			"Complement": complementDna,
			"RevComp":    revCompDna,
		},
		AlphabetIs(alphabet.Dna),
	)
	return n, n.Validate()
}

// NewDnaIupac generates a new sequence that validates against the DnaIupac alphabet
func NewDnaIupac(s string) (*Sequence, error) {
	n := New(
		s,
		map[string]SeqFunc{
			"Reverse":    reverseDnaIupac,
			"Complement": complementDnaIupac,
			"RevComp":    revCompDnaIupac,
		},
		AlphabetIs(alphabet.DnaIupac),
	)
	return n, n.Validate()
}

// NewRna generates a new sequence that validates against the Rna alphabet
func NewRna(s string) (*Sequence, error) {
	n := New(
		s,
		map[string]SeqFunc{
			"Reverse":    reverseRna,
			"Complement": complementRna,
			"RevComp":    revCompRna,
		},
		AlphabetIs(alphabet.Rna),
	)
	return n, n.Validate()
}

// NewRnaIupac generates a new sequence that validates against the RnaIupac alphabet
func NewRnaIupac(s string) (*Sequence, error) {
	n := New(
		s,
		map[string]SeqFunc{
			"Reverse":    reverseRnaIupac,
			"Complement": complementRnaIupac,
			"RevComp":    revCompRnaIupac,
		},
		AlphabetIs(alphabet.RnaIupac),
	)
	return n, n.Validate()
}
