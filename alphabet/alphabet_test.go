package alphabet

import (
	"fmt"
	"testing"
)

func TestDna(t *testing.T) {
	var a Interface = Dna
	t.Run("Correct length", IsExpectedLength(a, 4))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "ATGC"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}

func TestDnaIupac(t *testing.T) {
	var a Interface = DnaIupac
	t.Run("Correct length", IsExpectedLength(a, 16))
	t.Run("Has gap", HasExpectedLetter(a, '-'))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "ATGC" + "RYSWKM" + "BDHVN"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}

func TestRna(t *testing.T) {
	var a Interface = Rna
	t.Run("Correct length", IsExpectedLength(a, 4))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "AUGC"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}

func TestRnaIupac(t *testing.T) {
	var a Interface = RnaIupac
	t.Run("Correct length", IsExpectedLength(a, 16))
	t.Run("Has gap", HasExpectedLetter(a, '-'))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "AUGC" + "RYSWKM" + "BDHVN"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}

func TestProtein(t *testing.T) {
	var a Interface = Protein
	t.Run("Correct length", IsExpectedLength(a, 20))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "ACDEFGHIKLMNPQRSTVWY"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}

func TestProteinGapped(t *testing.T) {
	var a Interface = ProteinGapped
	t.Run("Correct length", IsExpectedLength(a, 21))
	t.Run("Has gap", HasExpectedLetter(a, '-'))
	t.Run("Expected letters", func(t *testing.T) {
		letters := "ACDEFGHIKLMNPQRSTVWY"
		for i, v := range a.Contains([]byte(letters)) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
}
