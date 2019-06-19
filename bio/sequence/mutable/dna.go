package mutable

import (
	"fmt"
	"strings"

	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/alphabet/hashmap"
	"github.com/rhagenson/bio-go/bio/data/codon"
	"github.com/rhagenson/bio-go/bio/sequence"
	"github.com/rhagenson/bio-go/bio/utils"
)

var _ sequence.Interface = new(Dna)
var _ sequence.Reverser = new(Dna)
var _ sequence.RevComper = new(Dna)
var _ sequence.Complementer = new(Dna)
var _ sequence.Transcriber = new(Dna)
var _ sequence.Translater = new(Dna)
var _ sequence.Alphabeter = new(Dna)
var _ sequence.LetterCounter = new(Dna)
var _ sequence.Validator = new(Dna)
var _ Wither = new(Dna)

// Dna is a sequence witch validates against the Dna alphabet
// and knows how to reverse, complement, and revcomp itself
type Dna struct {
	*Struct
}

// NewDna generates a New sequence that validates against the Dna alphabet
func NewDna(s string) (*Dna, error) {
	n := New(
		s,
		sequence.AlphabetIs(hashmap.NewDna()),
	)
	return &Dna{n}, n.Validate()
}

// Reverse is the same Dna with the sequence reversed
func (x *Dna) Reverse() (sequence.Interface, error) {
	l := x.Length()
	t := make([]string, l)
	var pos1, pos2 string
	for i := uint(0); i < l/2; i++ {
		pos1, _ = x.Position(i)
		pos2, _ = x.Position(l - 1 - i)
		t[i], t[l-1-i] = pos2, pos1
	}
	x.seq = strings.Join(t, "")
	return x, x.Validate()
}

// RevComp is the same Dna with the sequence reversed and complemented
func (x *Dna) RevComp() (sequence.Interface, error) {
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := make([]string, l)
	var pos1, pos2 string
	for i := uint(0); i < l/2; i++ {
		pos1, _ = x.Position(i)
		pos2, _ = x.Position(l - 1 - i)
		t[i], t[l-1-i] = c.Complement(pos2), c.Complement(pos1)
	}
	x.seq = strings.Join(t, "")
	return x, x.Validate()
}

// Complement is the same Dna with the sequence complemented
func (x *Dna) Complement() (sequence.Interface, error) {
	c := x.Alphabet().(alphabet.Complementer)
	l := x.Length()
	t := make([]string, l)
	var pos string
	for i := uint(0); i < l/2; i++ {
		pos, _ = x.Position(i)
		t[i] = c.Complement(pos)
	}
	x.seq = strings.Join(t, "")
	return x, x.Validate()
}

// Transcribe returns the DNA->RNA transcription product
func (x *Dna) Transcribe() (sequence.Interface, error) {
	t := []byte(x.String())
	for i, c := range t {
		if c == 'T' {
			t[i] = 'U'
		} else if c == 't' {
			t[i] = 'u'
		}
	}
	return NewRna(string(t))
}

// Translate returns a translated genetic product made from using a codon table
// The stop argument determines which character to use for indicating a stop codon
// If any stop codon is found and the stop argument is not a valid character in the
// Protein alphabet, an error will result stating as such. Therefore, if a stop
// codon is expected, checking the error message for the quoted character should be done.
func (x *Dna) Translate(table codon.Interface, stop byte) (sequence.Interface, error) {
	seq := x.String()
	t := make([]byte, len(seq)/3)

	var ok bool
	for i := range t {
		cdn := seq[i*3 : i*3+3]

		if utils.InStrings(cdn, table.StopCodons()) {
			t[i] = stop
		} else {
			t[i], ok = table.Translate(cdn)
			if !ok {
				return nil, fmt.Errorf("failed to translate codon: %q when using %s", cdn, table)
			}
		}
	}

	return NewProtein(string(t))
}

// Alphabet reveals the underlying alphabet in use
func (x *Dna) Alphabet() alphabet.Interface {
	return hashmap.NewDna()
}

// LetterCount reveals the number of occurrences for each letter in a sequence
func (x *Dna) LetterCount() map[string]uint {
	return sequence.LetterCount(x)
}
