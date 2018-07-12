package alphabet

import "testing"

var _ Interface = new(Protein)

func TestProtein(t *testing.T) {
	a := new(Protein)
	t.Run("Correct length", IsExpectedLength(a, 20))
}
