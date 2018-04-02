package letter

import (
	"fmt"
	"testing"
)

func TestLetterIsStringer(t *testing.T) {
	var _ fmt.Stringer = Letter("")
}
