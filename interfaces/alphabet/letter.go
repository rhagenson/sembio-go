package alphabet

import "fmt"

// Letter is an arbitrary entry in an Alphabet (string allows multirune entries in the Alphabet)
type Letter string

var _ fmt.Stringer = Letter("")

func (l Letter) String() string { return string(l) }
