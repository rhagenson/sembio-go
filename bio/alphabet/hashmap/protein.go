package hashmap

import "github.com/rhagenson/bio-go/bio/alphabet"

// Protein is the twenty letter standard encoding
type Protein struct {
	*Struct
}

func NewProtein() *Protein {
	return &Protein{
		New(alphabet.ProteinLetters),
	}
}
