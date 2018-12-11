package sequence

import (
	"fmt"

	"bitbucket.org/rhagenson/bio/alphabet"
)

// ValFunc is a function that validates a sequence
type ValFunc func(*backer) error

// Validator provides a variadic method to validate the sequence
type Validator interface {
	Validate() error
}

// AlphabetIs specifies whether a sequence conforms to a given Alphabet
func AlphabetIs(a *alphabet.Alphabet) ValFunc {
	return ValFunc(
		func(x *backer) error {
			for i := uint(0); i < x.Length()+a.Width()-1; i = i + a.Width() {
				letter := x.seq[i : i+a.Width()]
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
