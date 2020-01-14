package mutable_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sembio/go/bio/alphabet"
	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/sequence"
	"github.com/sembio/go/bio/sequence/mutable"
	"github.com/sembio/go/bio/test"
	"github.com/sembio/go/bio/utils"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedProteinGapped(t *testing.T) {
	s, _ := mutable.NewProteinGapped("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestProteinGappedHasMethods(t *testing.T) {
	s, _ := mutable.NewProteinGapped("")

	t.Run("Has Reverse method", func(t *testing.T) {
		if _, err := s.Reverse(); err != nil {
			t.Error("Reverse method does not exist")
		}
	})
}

func TestProteinGappedCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ProteinGapped is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				seq, _ := mutable.NewProteinGapped(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("ProteinGapped has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				seq, _ := mutable.NewProteinGapped(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("ProteinGapped has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				seq, _ := mutable.NewProteinGapped(s)
				onefourth := n / 4
				threefourths := n * 3 / 4
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("ProteinGapped has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				seq, _ := mutable.NewProteinGapped(s)
				onefourth := n / 4
				threefourth := n * (3 / 4)
				gotoneforth, _ := seq.Position(onefourth)
				wantoneforth := string(s[onefourth])
				gotthreeforth, _ := seq.Position(threefourth)
				wantthreeforth := string(s[threefourth])
				return gotoneforth == wantoneforth && gotthreeforth == wantthreeforth
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinGappedMutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				original, _ := mutable.NewProteinGapped(s)
				clone, _ := mutable.NewProteinGapped(s)
				original.With(mutable.PositionAs(n/2, t))
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("WithRange mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				original, _ := mutable.NewProteinGapped(s)
				clone, _ := mutable.NewProteinGapped(s)
				original.With(mutable.RangeAs(n/4, n*3/4, t))
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Reverse mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				if s == utils.ReverseByBytes(s) { // Skip palindromes
					return true
				}
				original, _ := mutable.NewProteinGapped(s)
				clone, _ := mutable.NewProteinGapped(s)
				original.Reverse()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinGappedMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				want, _ := mutable.NewProteinGapped(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*mutable.ProteinGapped).Reverse()
				return want.String() == got.(*mutable.ProteinGapped).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinGappedErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(string(alphabet.TestExcludesSingleLetters([]byte(hashmap.NewProteinGapped().String())))),
				)
				if _, err := mutable.NewProteinGapped(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("ProteinGapped creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("ProteinGapped should error when using invalid characters, error")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("start > stop errors",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				seq, _ := mutable.NewProteinGapped(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("ProteinGapped should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("ProteinGapped Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinGappedParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("mutable.NewProteinGapped(s) == mutable.NewProteinGapped(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				ret := make(chan *mutable.ProteinGapped)
				go func(s string, out chan *mutable.ProteinGapped) {
					seq, _ := mutable.NewProteinGapped(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *mutable.ProteinGapped) {
					seq, _ := mutable.NewProteinGapped(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.(*mutable.ProteinGapped).Reverse() == seq.(*mutable.ProteinGapped).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				ret := make(chan *mutable.ProteinGapped)
				seq, _ := mutable.NewProteinGapped(s)
				go func(seq *mutable.ProteinGapped, out chan *mutable.ProteinGapped) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.ProteinGapped)
				}(seq, ret)
				go func(seq *mutable.ProteinGapped, out chan *mutable.ProteinGapped) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.ProteinGapped)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

// Building a new ProteinGapped from valid letters results in no error
func ExampleNewProteinGapped_errorless() {
	s, err := mutable.NewProteinGapped("ACDEFGHIKLMNPQRSTVWY" + "-")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// ACDEFGHIKLMNPQRSTVWY-, <nil>
}

// Building a new ProteinGapped from invalid letters results in an error
// Note that only the first error is returned, not all errors
// The invalid '%' is caught, but nothing is said of the invalid '&'
func ExampleNewProteinGapped_errored() {
	s, err := mutable.NewProteinGapped("%" + "ACDEFGHIKLMNPQRSTVWY" + "-" + "&")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// %ACDEFGHIKLMNPQRSTVWY-&, "%" not in alphabet
}

// Reversing a valid ProteinGapped results in no error
func ExampleProteinGapped_Reverse_errorless() {
	s, _ := mutable.NewProteinGapped("ACDEFGHIKLMNPQRSTVWY" + "-")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// -YWVTSRQPNMLKIHGFEDCA, <nil>
}

// Reversing an invalid ProteinGapped results in an error
// Note that only the first error is returned, not all errors
// The invalid '&' is caught, but nothing is said of the invalid '%'
func ExampleProteinGapped_Reverse_errored() {
	s, _ := mutable.NewProteinGapped("%" + "ACDEFGHIKLMNPQRSTVWY" + "-" + "&")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// &-YWVTSRQPNMLKIHGFEDCA%, "&" not in alphabet
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewProteinGapped()
func ExampleProteinGapped_Alphabet() {
	s, _ := mutable.NewProteinGapped("ACDEFGHIKLMNPQRSTVWY" + "-")

	fmt.Println(s.Alphabet())
	// Output:
	// -ACDEFGHIKLMNPQRSTVWY
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewProteinGapped()
func ExampleProteinGapped_LetterCount() {
	s, _ := mutable.NewProteinGapped("ACDEFGHIKLMNPQRSTVWY" + "-" + "NNNN")

	fmt.Println(s.LetterCount())
	// Output:
	// map[-:1 A:1 C:1 D:1 E:1 F:1 G:1 H:1 I:1 K:1 L:1 M:1 N:5 P:1 Q:1 R:1 S:1 T:1 V:1 W:1 Y:1]
}
