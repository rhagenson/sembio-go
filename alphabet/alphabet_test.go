package alphabet_test

import (
	"bytes"
	"fmt"
	"testing"

	"bitbucket.org/rhagenson/bio/alphabet"
)

func TestTesting(t *testing.T) {
	t.Run("IsExpectedLength", func(t *testing.T) {
		a := alphabet.New("N")
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
		a := alphabet.New("N")
		t.Run("Success", alphabet.HasExpectedLetter(a, 'N'))
		t.Run("Failure", func(t *testing.T) {
			var t2 = new(testing.T)
			alphabet.HasExpectedLetter(a, 'X')(t2)
			if !t2.Failed() {
				t.Errorf("Failure case incorrectly passed.")
			}
		})
	})
	t.Run("NotLetters", func(t *testing.T) {
		letter := byte('A') // TODO: Replace with a random ASCII character
		t.Run("Success", func(t *testing.T) {
			if bytes.IndexByte(alphabet.NotLetters([]byte{letter}), letter) != -1 {
				t.Errorf(fmt.Sprintf("%q was found when it should not have been", letter))
			}
		})
		t.Run("Failure", func(t *testing.T) {
			nextLetter := letter
			if letter += 1; letter > 'z' {
				nextLetter = 'a'
			}
			if bytes.IndexByte(alphabet.NotLetters([]byte{letter}), nextLetter) == 1 {
				t.Errorf("Failure case incorrectly passed.")
			}
		})
	})
}

func TestDna(t *testing.T) {
	var a alphabet.Interface = alphabet.Dna
	letters := []byte("ATGC")
	notLetters := alphabet.NotLetters(letters)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 4))
	t.Run("Expected letters", func(t *testing.T) {
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		for i, v := range a.Contains(notLetters...) {
			t.Run(fmt.Sprintf("Excludes %q", notLetters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", notLetters[i])
				}
			})
		}
	})
}

func TestDnaIupac(t *testing.T) {
	var a alphabet.Interface = alphabet.DnaIupac
	letters := []byte("ATGC" + "RYSWKM" + "BDHVN")
	notLetters := alphabet.NotLetters(letters)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, '-'))
	t.Run("Expected letters", func(t *testing.T) {
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		for i, v := range a.Contains(notLetters...) {
			t.Run(fmt.Sprintf("Excludes %q", notLetters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", notLetters[i])
				}
			})
		}
	})
}

func TestRna(t *testing.T) {
	var a alphabet.Interface = alphabet.Rna
	letters := []byte("AUGC")
	notLetters := alphabet.NotLetters(letters)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 4))
	t.Run("Expected letters", func(t *testing.T) {
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		for i, v := range a.Contains(notLetters...) {
			t.Run(fmt.Sprintf("Excludes %q", notLetters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", notLetters[i])
				}
			})
		}
	})
}

func TestRnaIupac(t *testing.T) {
	var a alphabet.Interface = alphabet.RnaIupac
	letters := []byte("AUGC" + "RYSWKM" + "BDHVN")
	notLetters := alphabet.NotLetters(letters)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, '-'))
	t.Run("Expected letters", func(t *testing.T) {
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		for i, v := range a.Contains(notLetters...) {
			t.Run(fmt.Sprintf("Excludes %q", notLetters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", notLetters[i])
				}
			})
		}
	})
}

func TestProtein(t *testing.T) {
	var a alphabet.Interface = alphabet.Protein
	letters := []byte("ACDEFGHIKLMNPQRSTVWY")
	notLetters := alphabet.NotLetters(letters)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 20))
	t.Run("Expected letters", func(t *testing.T) {
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		for i, v := range a.Contains(notLetters...) {
			t.Run(fmt.Sprintf("Excludes %q", notLetters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", notLetters[i])
				}
			})
		}
	})
}

func TestProteinGapped(t *testing.T) {
	var a alphabet.Interface = alphabet.ProteinGapped
	letters := []byte("ACDEFGHIKLMNPQRSTVWY")
	notLetters := alphabet.NotLetters(letters)
	t.Run("Correct length", alphabet.IsExpectedLength(a, 21))
	t.Run("Has gap", alphabet.HasExpectedLetter(a, '-'))
	t.Run("Expected letters", func(t *testing.T) {
		for i, v := range a.Contains(letters...) {
			t.Run(fmt.Sprintf("Contains %q", letters[i]), func(t *testing.T) {
				if !v {
					t.Errorf("Does not contain %q", letters[i])
				}
			})
		}
	})
	t.Run("Incorrect letters", func(t *testing.T) {
		for i, v := range a.Contains(notLetters...) {
			t.Run(fmt.Sprintf("Excludes %q", notLetters[i]), func(t *testing.T) {
				if v {
					t.Errorf("Should not contain %q", notLetters[i])
				}
			})
		}
	})
}
