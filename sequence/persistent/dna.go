package persistent

import (
	"fmt"

	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/data/codon"
	"bitbucket.org/rhagenson/bio/helpers/complement"
	"bitbucket.org/rhagenson/bio/sequence"
)

var _ sequence.Reverser = new(Dna)
var _ sequence.RevComper = new(Dna)
var _ sequence.Complementer = new(Dna)
var _ sequence.Transcriber = new(Dna)
var _ sequence.Translater = new(Dna)
var _ Wither = new(Dna)

// Dna is a sequence witch validates against the Dna alphabet
// and knows how to reverse, complement, and revcomp itself
type Dna struct {
	*Struct
}

// NewDna generates a New sequence that validates against the Dna alphabet
func NewDna(s string) (*Dna, error) {
	n := NewStruct(
		s,
		AlphabetIs(alphabet.Dna),
	)
	return &Dna{n}, n.Validate()
}

// Reverse is the same Dna with the sequence reversed
func (x *Dna) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewDna(string(t))
}

// RevComp is the same Dna with the sequence reversed and complemented
func (x *Dna) RevComp() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.Dna(t[l-1-i]), complement.Dna(t[i])
	}
	return NewDna(string(t))
}

// Complement is the same DnaIupac with the sequence complemented
func (x *Dna) Complement() (sequence.Interface, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.Dna(byte(x.seq[i]))
	}
	return NewDna(string(t))
}

func (x *Dna) Transcribe() (sequence.Interface, error) {
	t := []byte(x.String())
	for i, c := range t {
		if c == 'T' {
			t[i] = 'U'
		}
	}
	return NewRna(string(t))
}

func (x *Dna) Translate(l codon.Translater) (sequence.Interface, error) {
	seq := x.String()
	t := make([]byte, len(seq)/3)

	var ok bool
	for i := range t {
		cdn := seq[i*3 : i*3+3]
		t[i], ok = l.Translate(cdn)
		if !ok {
			return nil, fmt.Errorf("failed to translate codon: %q when using %s", cdn, l)
		}
	}
	return NewProtein(string(t))
}
