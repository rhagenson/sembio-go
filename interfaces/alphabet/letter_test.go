package alphabet

import (
	"fmt"
	"testing"
)

func TestLetterIsStringer(t *testing.T) {
	var _ fmt.Stringer = Letter("")
}
