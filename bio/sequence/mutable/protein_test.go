package mutable_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/sequence"
	"github.com/bio-ext/bio-go/bio/sequence/mutable"
	"github.com/bio-ext/bio-go/bio/test"
	"github.com/bio-ext/bio-go/bio/utils"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedProtein(t *testing.T) {
	s, _ := mutable.NewProtein("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestProteinHasMethods(t *testing.T) {
	s, _ := mutable.NewProtein("")

	t.Run("Has Reverse method", func(t *testing.T) {
		if _, err := s.Reverse(); err != nil {
			t.Error("Reverse method does not exist")
		}
	})
}

func TestProteinCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Protein is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				seq, _ := mutable.NewProtein(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Protein has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				seq, _ := mutable.NewProtein(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Protein has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				seq, _ := mutable.NewProtein(s)
				onefourth := n / 4
				threefourths := n * 3 / 4
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Protein has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				seq, _ := mutable.NewProtein(s)
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

func TestProteinMutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				original, _ := mutable.NewProtein(s)
				clone, _ := mutable.NewProtein(s)
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
					[]rune(hashmap.NewProtein().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				original, _ := mutable.NewProtein(s)
				clone, _ := mutable.NewProtein(s)
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
					[]rune(hashmap.NewProtein().String()),
				)
				if s == utils.ReverseByBytes(s) { // Skip palindromes
					return true
				}
				original, _ := mutable.NewProtein(s)
				clone, _ := mutable.NewProtein(s)
				original.Reverse()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				want, _ := mutable.NewProtein(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*mutable.Protein).Reverse()
				return want.String() == got.(*mutable.Protein).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(string(alphabet.TestExcludesSingleLetters([]byte(hashmap.NewProtein().String())))),
				)
				if _, err := mutable.NewProtein(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("Protein creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("Protein should error when using invalid characters, error")
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
					[]rune(hashmap.NewProtein().String()),
				)
				seq, _ := mutable.NewProtein(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("Protein should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("Protein Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("mutable.NewProtein(s) == mutable.NewProtein(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				ret := make(chan *mutable.Protein)
				go func(s string, out chan *mutable.Protein) {
					seq, _ := mutable.NewProtein(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *mutable.Protein) {
					seq, _ := mutable.NewProtein(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.(*mutable.Protein).Reverse() == seq.(*mutable.Protein).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				ret := make(chan *mutable.Protein)
				seq, _ := mutable.NewProtein(s)
				go func(seq *mutable.Protein, out chan *mutable.Protein) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.Protein)
				}(seq, ret)
				go func(seq *mutable.Protein, out chan *mutable.Protein) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.Protein)
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

// Building a new Protein from valid letters results in no error
func ExampleNewProtein_errorless() {
	s, err := mutable.NewProtein("ACDEFGHIKLMNPQRSTVWY")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// ACDEFGHIKLMNPQRSTVWY, <nil>
}

// Building a new Protein from invalid letters results in an error
// Note that only the first error is returned, not all errors
// The invalid '%' is caught, but nothing is said of the invalid '&'
func ExampleNewProtein_errored() {
	s, err := mutable.NewProtein("%" + "ACDEFGHIKLMNPQRSTVWY" + "&")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// %ACDEFGHIKLMNPQRSTVWY&, "%" not in alphabet
}

// Reversing a valid Protein results in no error
func ExampleProtein_Reverse_errorless() {
	s, _ := mutable.NewProtein("ACDEFGHIKLMNPQRSTVWY")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// YWVTSRQPNMLKIHGFEDCA, <nil>
}

// Reversing an invalid Protein results in an error
// Note that only the first error is returned, not all errors
// The invalid '&' is caught, but nothing is said of the invalid '%'
func ExampleProtein_Reverse_errored() {
	s, _ := mutable.NewProtein("%" + "ACDEFGHIKLMNPQRSTVWY" + "&")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// &YWVTSRQPNMLKIHGFEDCA%, "&" not in alphabet
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewProtein()
func ExampleProtein_Alphabet() {
	s, _ := mutable.NewProtein("ACDEFGHIKLMNPQRSTVWY")

	fmt.Println(s.Alphabet())
	// Output:
	// ACDEFGHIKLMNPQRSTVWY
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewProtein()
func ExampleProtein_LetterCount() {
	s, _ := mutable.NewProtein("ACDEFGHIKLMNPQRSTVWY" + "NNNN")

	fmt.Println(s.LetterCount())
	// Output:
	// map[A:1 C:1 D:1 E:1 F:1 G:1 H:1 I:1 K:1 L:1 M:1 N:5 P:1 Q:1 R:1 S:1 T:1 V:1 W:1 Y:1]
}
