package immutable

import (
	"fmt"

	"bitbucket.org/rhagenson/bio/alphabet"
)

// ValFunc is a function that validates a sequence
type ValFunc func(*Struct) error

// Validator provides a variadic method to validate the sequence
type Validator interface {
	Validate() error
}

// AlphabetIs specifies whether a sequence conforms to a given Alphabet
func AlphabetIs(a alphabet.Interface) ValFunc {
	return ValFunc(
		func(x *Struct) error {
			for i := uint(0); i < x.Length(); i++ {
				letter := x.seq[i]
				for _, found := range a.Contains(letter) {
					if !found {
						return fmt.Errorf("%q not in alphabet", letter)
					}
				}
			}
			return nil
		},
	)
}
