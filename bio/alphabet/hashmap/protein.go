package hashmap

import "github.com/sembio/go/bio/alphabet"

// Protein is the twenty letter standard encoding
type Protein struct {
	*Struct
}

// NewProtein generates a protein alphabet
func NewProtein() *Protein {
	return &Protein{
		New(alphabet.ProteinLetters),
	}
}
