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
		t.Run("Failure", func(t *testing.T) {
			var t2 = new(testing.T)
			alphabet.IsExpectedLength(a, 2)(t2)
			if !t2.Failed() {
				t.Errorf("Failure case incorrectly passed.")
			}
		})
	})
	t.Run("HasExpectedLetter", func(t *testing.T) {
		a := alphabet.New("N", 1)
		t.Run("Success", alphabet.HasExpectedLetter(a, "N"))
		t.Run("Failure", func(t *testing.T) {
			var t2 = new(testing.T)
			alphabet.HasExpectedLetter(a, "X")(t2)
			if !t2.Failed() {
				t.Errorf("Failure case incorrectly passed.")
			}
		})
	})
	t.Run("TestSplitByN", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			exp := []string{"A", "T", "G", "C"}
			got := alphabet.TestSplitByN("ATGC", 1)
			for i := range exp {
				if got[i] != exp[i] {
					t.Errorf("Got %q, expected %q.", got[i], exp[i])
				}
			}
		})
		t.Run("Failure", func(t *testing.T) {
			got := alphabet.TestSplitByN("ATGC", 0)
			if got != nil {
				t.Errorf("Got %q, expected nil.", got)
			}
		})
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
