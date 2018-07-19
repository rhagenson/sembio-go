package sequence

import (
	"fmt"

	"bitbucket.org/rhagenson/bigr/alphabet"
)

// ValFunc is a function that validates a sequence
type ValFunc func(*Sequence) error

// Validator provides a variadic method to validate the sequence
type Validator interface {
	Validate() error
}

// AlphabetIs specifies whether a Sequence conforms to a given Alphabet
func AlphabetIs(a alphabet.Alphabet) ValFunc {
	return ValFunc(
		func(x *Sequence) error {
			for _, l := range x.seq {
				for _, b := range a.Contains(byte(l)) {
					if !b {
						return fmt.Errorf("%q not in alphabet", l)
					}
				}
			}
			return nil
		},
	)
}
