package alphabet

import (
	"testing"
)

var _ Interface = new(ProteinGapped)

func TestProteinGapped(t *testing.T) {
	a := new(ProteinGapped)
	t.Run("Correct length", IsExpectedLength(a, 21))
	t.Run("Has gap", HasExpectedLetter(a, '-'))
}
