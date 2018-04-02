package alphabet

// Gapper implements the Gapped method, which tells a user if a struct contains a gap Letter
type Gapper interface {
	// Gapped tells the user if the struct includes a gap character of some kind
	Gapped() bool
}
