package alphabet_test

import (
	"fmt"
	"testing"

	"bitbucket.org/rhagenson/bio/alphabet"
)

func TestDna(t *testing.T) {
	var a alphabet.Interface = alphabet.Dna
	t.Run("Correct length", alphabet.IsExpectedLength(a, 4))
	t.Run("Expected letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("ATGC", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("XJZ", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Excludes %q", letters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", letters[i])
				}
			})
		}
	})
}

func TestDnaIupac(t *testing.T) {
	var a alphabet.Interface = alphabet.DnaIupac
	t.Run("Correct length", alphabet.IsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, "-"))
	t.Run("Expected letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("ATGC"+"RYSWKM"+"BDHVN", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("XJZ", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Excludes %q", letters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", letters[i])
				}
			})
		}
	})
}

func TestRna(t *testing.T) {
	var a alphabet.Interface = alphabet.Rna
	t.Run("Correct length", alphabet.IsExpectedLength(a, 4))
	t.Run("Expected letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("AUGC", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("XJZ", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Excludes %q", letters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", letters[i])
				}
			})
		}
	})
}

func TestRnaIupac(t *testing.T) {
	var a alphabet.Interface = alphabet.RnaIupac
	t.Run("Correct length", alphabet.IsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, "-"))
	t.Run("Expected letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("AUGC"+"RYSWKM"+"BDHVN", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("XJZ", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Excludes %q", letters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", letters[i])
				}
			})
		}
	})
}

func TestProtein(t *testing.T) {
	var a alphabet.Interface = alphabet.Protein
	t.Run("Correct length", alphabet.IsExpectedLength(a, 20))
	t.Run("Expected letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("ACDEFGHIKLMNPQRSTVWY", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("XJZ", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Excludes %q", letters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", letters[i])
				}
			})
		}
	})
}

func TestProteinGapped(t *testing.T) {
	var a alphabet.Interface = alphabet.ProteinGapped
	t.Run("Correct length", alphabet.IsExpectedLength(a, 21))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, "-"))
	t.Run("Expected letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("ACDEFGHIKLMNPQRSTVWY", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		letters := alphabet.TestSplitByN("XJZ", 1)
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Excludes %q", letters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", letters[i])
				}
			})
		}
	})
}
