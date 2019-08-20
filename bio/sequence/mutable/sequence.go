package mutable

import (
	"fmt"

	"github.com/bio-ext/bio-go/bio/sequence"
)

var _ sequence.Interface = new(Struct)
var _ Wither = new(Struct)
var _ sequence.Validator = new(Struct)

// Struct stores a linear sequence and has optional validators
type Struct struct {
	seq        string
	validators []sequence.ValFunc
}

// New generates a new generalized sequence with optional validators
func New(s string, vs ...sequence.ValFunc) *Struct {
	seq := &Struct{
		seq:        s,
		validators: make([]sequence.ValFunc, 0),
	}
	seq.validators = append(seq.validators, vs...)
	return seq
}

// With runs a series of transformative actions, returning the final result
// Attention: With does not call Validate.
func (x *Struct) With(fs ...WithFunc) *Struct {
	for _, f := range fs {
		x = f(x)
	}
	return x
}

// Validate runs a series of Validator funcs, returning the first error
func (x *Struct) Validate() error {
	for _, f := range x.validators {
		err := f(x)
		if err != nil {
			return err
		}
	}
	return nil
}

// Length is the number of positions in the sequence
func (x *Struct) Length() uint {
	return uint(len(x.seq))
}

// Position is the letter found at position n
func (x *Struct) Position(n uint) (string, error) {
	if n < x.Length() {
		return string(x.seq[n]), nil
	}
	return "", fmt.Errorf("requested impossible position [%d]", n)
}

// Range is the letters found in the half-open range
func (x *Struct) Range(st, sp uint) (string, error) {
	if sp == x.Length() {
		return x.seq[st:], nil
	} else if st < sp && sp < x.Length() {
		return x.seq[st:sp], nil
	}
	return "", fmt.Errorf("requested impossible range [%d:%d]", st, sp)
}

// String reveals the underlying string
func (x *Struct) String() string {
	return x.seq
}
