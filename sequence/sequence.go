package sequence

import (
	"fmt"
)

// backer stores a linear sequence and has optional validators
type backer struct {
	seq        string
	validators []ValFunc
}

// newBacker generates a new backing sequence with optional validators
func newBacker(s string, vs ...ValFunc) *backer {
	seq := &backer{
		seq:        s,
		validators: make([]ValFunc, 0),
	}
	seq.validators = append(seq.validators, vs...)
	return seq
}

// With runs a series of transformative actions, returning the final result
// Attention: With does not call Validate.
func (x *backer) With(fs ...WithFunc) *backer {
	y := new(backer)
	*y = *x // Create copy to protext the receiver from Wither funcs
	for _, f := range fs {
		y = f(y)
	}
	return y
}

// Validate runs a series of Validator funcs, returning the first error
func (x *backer) Validate() error {
	y := new(backer)
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
func (x *backer) Length() uint {
	return uint(len(x.seq))
}

// Position is the letter found at position n
func (x *backer) Position(n uint) (string, error) {
	if n < x.Length() {
		return string(x.seq[n]), nil
	}
	return "", fmt.Errorf("requested impossible position [%d]", n)
}

// Range is the letters found in the half-open range
func (x *backer) Range(st, sp uint) (string, error) {
	if sp == x.Length() {
		return x.seq[st:], nil
	} else if st < sp && sp < x.Length() {
		return x.seq[st:sp], nil
	}
	return "", fmt.Errorf("requested impossible range [%d:%d]", st, sp)
}
