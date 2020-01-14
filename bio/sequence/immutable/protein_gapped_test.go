package immutable_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sembio/go/bio/alphabet"
	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/sequence"
	"github.com/sembio/go/bio/sequence/immutable"
	"github.com/sembio/go/bio/test"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedProteinGapped(t *testing.T) {
	s, _ := immutable.NewProteinGapped("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestProteinGappedHasMethods(t *testing.T) {
	s, _ := immutable.NewProteinGapped("")

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
				seq, _ := immutable.NewProteinGapped(s)
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
				seq, _ := immutable.NewProteinGapped(s)
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
				seq, _ := immutable.NewProteinGapped(s)
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
				seq, _ := immutable.NewProteinGapped(s)
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

func TestProteinGappedImmutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
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
				original, _ := immutable.NewProteinGapped(s)
				clone, _ := immutable.NewProteinGapped(s)
				original.With(immutable.PositionAs(n*(1/2), t))
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("WithRange does not mutate in-place",
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
				original, _ := immutable.NewProteinGapped(s)
				clone, _ := immutable.NewProteinGapped(s)
				original.With(immutable.RangeAs(n*(1/4), n*(3/4), t))
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Reverse does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				original, _ := immutable.NewProteinGapped(s)
				clone, _ := immutable.NewProteinGapped(s)
				original.Reverse()
				return original.String() == clone.String()
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
				want, _ := immutable.NewProteinGapped(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*immutable.ProteinGapped).Reverse()
				return want.String() == got.(*immutable.ProteinGapped).String()
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
				if _, err := immutable.NewProteinGapped(s); err != nil {
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
				seq, _ := immutable.NewProteinGapped(s)
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

	properties.Property("immutable.NewProteinGapped(s) == immutable.NewProteinGapped(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				ret := make(chan *immutable.ProteinGapped)
				go func(s string, out chan *immutable.ProteinGapped) {
					seq, _ := immutable.NewProteinGapped(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *immutable.ProteinGapped) {
					seq, _ := immutable.NewProteinGapped(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.(*immutable.ProteinGapped).Reverse() == seq.(*immutable.ProteinGapped).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProteinGapped().String()),
				)
				ret := make(chan *immutable.ProteinGapped)
				seq, _ := immutable.NewProteinGapped(s)
				go func(seq *immutable.ProteinGapped, out chan *immutable.ProteinGapped) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.ProteinGapped)
				}(seq, ret)
				go func(seq *immutable.ProteinGapped, out chan *immutable.ProteinGapped) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.ProteinGapped)
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
	s, err := immutable.NewProteinGapped("ACDEFGHIKLMNPQRSTVWY" + "-")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// ACDEFGHIKLMNPQRSTVWY-, <nil>
}

// Building a new ProteinGapped from invalid letters results in an error
// Note that only the first error is returned, not all errors
// The invalid '%' is caught, but nothing is said of the invalid '&'
func ExampleNewProteinGapped_errored() {
	s, err := immutable.NewProteinGapped("%" + "ACDEFGHIKLMNPQRSTVWY" + "-" + "&")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// %ACDEFGHIKLMNPQRSTVWY-&, "%" not in alphabet
}

// Reversing a valid ProteinGapped results in no error
func ExampleProteinGapped_Reverse_errorless() {
	s, _ := immutable.NewProteinGapped("ACDEFGHIKLMNPQRSTVWY" + "-")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// -YWVTSRQPNMLKIHGFEDCA, <nil>
}

// Reversing an invalid ProteinGapped results in an error
// Note that only the first error is returned, not all errors
// The invalid '&' is caught, but nothing is said of the invalid '%'
func ExampleProteinGapped_Reverse_errored() {
	s, _ := immutable.NewProteinGapped("%" + "ACDEFGHIKLMNPQRSTVWY" + "-" + "&")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// &-YWVTSRQPNMLKIHGFEDCA%, "&" not in alphabet
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to immutable.NewProteinGapped()
func ExampleProteinGapped_Alphabet() {
	s, _ := immutable.NewProteinGapped("ACDEFGHIKLMNPQRSTVWY" + "-")

	fmt.Println(s.Alphabet())
	// Output:
	// -ACDEFGHIKLMNPQRSTVWY
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to immutable.NewProteinGapped()
func ExampleProteinGapped_LetterCount() {
	s, _ := immutable.NewProteinGapped("ACDEFGHIKLMNPQRSTVWY" + "-" + "NNNN")

	fmt.Println(s.LetterCount())
	// Output:
	// map[-:1 A:1 C:1 D:1 E:1 F:1 G:1 H:1 I:1 K:1 L:1 M:1 N:5 P:1 Q:1 R:1 S:1 T:1 V:1 W:1 Y:1]
}
