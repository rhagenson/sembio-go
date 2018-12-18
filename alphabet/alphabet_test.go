package alphabet_test

import (
	"fmt"
	"testing"

	"bitbucket.org/rhagenson/bio/alphabet"
)

func TestTesting(t *testing.T) {
	t.Run("IsExpectedLength", func(t *testing.T) {
		a := alphabet.New("N", 1)
		t.Run("Success", alphabet.IsExpectedLength(a, 1))
		t.Run("Failure", alphabet.IsExpectedLength(a, 2))
		// Wrapping the failing case in a new *testing.T
		// and checking for T.Failed() == true is the correct manner
		// Perhaps needs a new bio/testing.go global FailCase helper?
		// Perhaps there is a library already checking for test failure
		// does the assert lib do this?
	})
}

func TestAlphabet(t *testing.T) {
	t.Run("Zero width becomes width of one", func(t *testing.T) {
		if alphabet.New("", 0).Width() == 0 {
			t.Error("Width can never be zero")
		}
	})
}

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
