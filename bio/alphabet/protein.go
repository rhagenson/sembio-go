package alphabet

// Protein is the twenty letter standard encoding
type Protein struct {
	*Struct
}

func NewProtein() *Protein {
	return &Protein{
		New(ProteinLetters),
	}
}
