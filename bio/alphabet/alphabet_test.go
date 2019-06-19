package alphabet_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/rhagenson/bio-go/bio/alphabet"
)

func TestTesting(t *testing.T) {
	t.Run("IsExpectedLength", func(t *testing.T) {
		a := alphabet.New("N")
		t.Run("Success", alphabet.TestIsExpectedLength(a, 1))
		t.Run("Failure", func(t *testing.T) {
			var t2 = new(testing.T)
			alphabet.TestIsExpectedLength(a, 2)(t2)
			if !t2.Failed() {
				t.Errorf("Failure case incorrectly passed.")
			}
		})
	})
	t.Run("HasExpectedLetter", func(t *testing.T) {
		a := alphabet.New("N")
		t.Run("Success", alphabet.TestHasExpectedLetter(a, 'N'))
		t.Run("Failure", func(t *testing.T) {
			var t2 = new(testing.T)
			alphabet.TestHasExpectedLetter(a, 'X')(t2)
			if !t2.Failed() {
				t.Errorf("Failure case incorrectly passed.")
			}
		})
	})
	t.Run("NotLetters", func(t *testing.T) {
		// Chooses random ASCII letter in range from
		// minimum position (A) to maximum position (z)
		letter := byte(rand.Intn('z'-'A') + 'A')
		t.Run("Success", func(t *testing.T) {
			if bytes.IndexByte(alphabet.TestExcludesLetters([]byte{letter}), letter) != -1 {
				t.Errorf(fmt.Sprintf("%q was found when it should not have been", letter))
			}
		})
		t.Run("Failure", func(t *testing.T) {
			nextLetter := letter
			if nextLetter += 1; nextLetter > 'z' {
				nextLetter = 'A' // Wrap to ASCII start
			}
			if bytes.IndexByte(alphabet.TestExcludesLetters([]byte{letter}), nextLetter) == -1 {
				t.Errorf("Failure case incorrectly passed.")
			}
		})
	})
}

func TestDna(t *testing.T) {
	var a alphabet.Interface = alphabet.NewDna()
	letters := []byte("ATGC")
	notLetters := alphabet.TestExcludesLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 4))
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
	var a alphabet.Interface = alphabet.NewDnaIupac()
	letters := []byte("ATGC" + "RYSWKM" + "BDHVN")
	notLetters := alphabet.TestExcludesLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.TestHasExpectedLetter(a, '-'))
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
	var a alphabet.Interface = alphabet.NewRna()
	letters := []byte("AUGC")
	notLetters := alphabet.TestExcludesLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 4))
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
	var a alphabet.Interface = alphabet.NewRnaIupac()
	letters := []byte("AUGC" + "RYSWKM" + "BDHVN")
	notLetters := alphabet.TestExcludesLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 16))
	t.Run("Has gap", alphabet.TestHasExpectedLetter(a, '-'))
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
	var a alphabet.Interface = alphabet.NewProtein()
	letters := []byte("ACDEFGHIKLMNPQRSTVWY")
	notLetters := alphabet.TestExcludesLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 20))
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
	var a alphabet.Interface = alphabet.NewProteinGapped()
	letters := []byte("ACDEFGHIKLMNPQRSTVWY")
	notLetters := alphabet.TestExcludesLetters(letters)
	t.Run("Correct length", alphabet.TestIsExpectedLength(a, 21))
	t.Run("Has gap", alphabet.TestHasExpectedLetter(a, '-'))
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
