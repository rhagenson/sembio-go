package sequence

import (
	"fmt"
)

// Sequence stores a linear sequence, accumulates its errors, and has
// optional validators
type Sequence struct {
	seq        string
	ops        map[string]SeqFunc
	validators []ValFunc
}

// New generates a new sequence with optional validators
func New(s string, ops map[string]SeqFunc, vs ...ValFunc) *Sequence {
	seq := &Sequence{
		seq:        s,
		ops:        ops,
		validators: make([]ValFunc, 0),
	}
	seq.validators = append(seq.validators, vs...)
	return seq
}

// With runs a series of transformative actions, returning the final result
// Attention: With does not call Validate.
func (x *Sequence) With(fs ...WithFunc) *Sequence {
	y := new(Sequence)
	*y = *x // Create copy to protext the receiver from Wither funcs
	for _, f := range fs {
		y = f(y)
	}
	return y
}

// Validate runs a series of Validator funcs, returning the first error
func (x *Sequence) Validate() error {
	y := new(Sequence)
	for _, f := range x.validators {
		*y = *x // Create a copy to protect Validator funcs from each other
		err := f(y)
		if err != nil {
			return err
		}
	}
	return nil
}

// Length is the number of positions in the sequence
func (x *Sequence) Length() uint {
	return uint(len(x.seq))
}

// Position is the letter found at position n
func (x *Sequence) Position(n uint) (string, error) {
	if n < x.Length() {
		return string(x.seq[n]), nil
	}
	return "", fmt.Errorf("requested impossible position [%d]", n)
}

// Range is the letters found in the half-open range
func (x *Sequence) Range(st, sp uint) (string, error) {
	if sp == x.Length() {
		return x.seq[st:], nil
	} else if st < sp && sp < x.Length() {
		return x.seq[st:sp], nil
	}
	return "", fmt.Errorf("requested impossible range [%d:%d]", st, sp)
}

// Reverse calls to the sequencing reversal op
func (x *Sequence) Reverse() (*Sequence, error) {
	return x.Op("Reverse")
}

// Complement calls to the sequencing complement op
func (x *Sequence) Complement() (*Sequence, error) {
	return x.Op("Complement")
}

// RevComp calls to the sequencing reverse-complement op
func (x *Sequence) RevComp() (*Sequence, error) {
	return x.Op("RevComp")
}

// Translate calls to the sequencing reverse-complement op
func (x *Sequence) Translate() (*Sequence, error) {
	return x.Op("Translate")
}

// Transcribe calls to the sequencing reverse-complement op
func (x *Sequence) Transcribe() (*Sequence, error) {
	return x.Op("Transcribe")
}

// Op calls a transformative operation
func (x *Sequence) Op(s string) (*Sequence, error) {
	if f, ok := x.ops[s]; ok {
		return f(x)
	}
	return x, fmt.Errorf("%q not implemented", s)
}

// RegisterOps registers new ops, or overwrites old ops
func (x *Sequence) RegisterOps(m map[string]SeqFunc) {
	for key, op := range m {
		x.ops[key] = op
	}
	return
}
