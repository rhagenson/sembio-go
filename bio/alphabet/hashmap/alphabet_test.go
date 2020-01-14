package hashmap_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/sembio/go/bio/alphabet"
	"github.com/sembio/go/bio/alphabet/hashmap"
)

func TestStruct(t *testing.T) {
	t.Run("IsExpectedLength", func(t *testing.T) {
		a := hashmap.New("N")
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
		a := hashmap.New("N")
		t.Run("Success", alphabet.TestHasExpectedLetter(a, "N"))
		t.Run("Failure", func(t *testing.T) {
			var t2 = new(testing.T)
			alphabet.TestHasExpectedLetter(a, "X")(t2)
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
			if bytes.IndexByte(alphabet.TestExcludesSingleLetters([]byte{letter}), letter) != -1 {
				t.Errorf(fmt.Sprintf("%q was found when it should not have been", letter))
			}
		})
		t.Run("Failure", func(t *testing.T) {
			nextLetter := letter
			if nextLetter += 1; nextLetter > 'z' {
				nextLetter = 'A' // Wrap to ASCII start
			}
			if bytes.IndexByte(alphabet.TestExcludesSingleLetters([]byte{letter}), nextLetter) == -1 {
				t.Errorf("Failure case incorrectly passed.")
			}
		})
	})
}

func ExampleNew() {
	a := hashmap.New("QWERTY")
	fmt.Println(a)
	// Output:
	// EQRTWY
}

func ExampleStruct_Length() {
	a := hashmap.New("QWERTY")
	fmt.Println(a.Length())
	// Output:
	// 6
}

func ExampleStruct_Contains() {
	a := hashmap.New("QWERTY")
	fmt.Println(a.Contains([]string{"Q", "WERTY", "A"}...))
	// Output:
	// [true false false]
}
