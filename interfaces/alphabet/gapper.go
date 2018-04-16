package alphabet

// Gapper implements the Gapped method, which tells a user if a struct contains gaps
type Gapper interface {
	// Gapped tells the user if the struct includes gaps
	Gapped() bool
}
