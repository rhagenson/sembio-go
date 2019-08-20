package alphabet_example

import (
	"fmt"

	alphabet "github.com/bio-ext/bio-go/bio/alphabet/hashmap"
)

// ExampleCustomAlphabet shows how to create a new alphabet
func ExampleCustomAlphabet() {
	// An alphabet is made from its individual letters
	a := alphabet.New("QWERTY")

	// Alphabets can printed
	fmt.Println(a.String())

	// Alphabets can check for given elements
	fmt.Println(a.Contains([]string{"Q", "WERTY", "A"}...))
}
