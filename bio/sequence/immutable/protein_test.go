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

func TestInitializedProtein(t *testing.T) {
	s, _ := immutable.NewProtein("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestProteinHasMethods(t *testing.T) {
	s, _ := immutable.NewProtein("")

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
				seq, _ := immutable.NewProtein(s)
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
				seq, _ := immutable.NewProtein(s)
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
				seq, _ := immutable.NewProtein(s)
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
				seq, _ := immutable.NewProtein(s)
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

func TestProteinImmutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
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
				original, _ := immutable.NewProtein(s)
				clone, _ := immutable.NewProtein(s)
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
					[]rune(hashmap.NewProtein().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				original, _ := immutable.NewProtein(s)
				clone, _ := immutable.NewProtein(s)
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
					[]rune(hashmap.NewProtein().String()),
				)
				original, _ := immutable.NewProtein(s)
				clone, _ := immutable.NewProtein(s)
				original.Reverse()
				return original.String() == clone.String()
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
				want, _ := immutable.NewProtein(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*immutable.Protein).Reverse()
				return want.String() == got.(*immutable.Protein).String()
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
				if _, err := immutable.NewProtein(s); err != nil {
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
				seq, _ := immutable.NewProtein(s)
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

	properties.Property("immutable.NewProtein(s) == immutable.NewProtein(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				ret := make(chan *immutable.Protein)
				go func(s string, out chan *immutable.Protein) {
					seq, _ := immutable.NewProtein(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *immutable.Protein) {
					seq, _ := immutable.NewProtein(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.(*immutable.Protein).Reverse() == seq.(*immutable.Protein).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewProtein().String()),
				)
				ret := make(chan *immutable.Protein)
				seq, _ := immutable.NewProtein(s)
				go func(seq *immutable.Protein, out chan *immutable.Protein) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.Protein)
				}(seq, ret)
				go func(seq *immutable.Protein, out chan *immutable.Protein) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.Protein)
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
	s, err := immutable.NewProtein("ACDEFGHIKLMNPQRSTVWY")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// ACDEFGHIKLMNPQRSTVWY, <nil>
}

// Building a new Protein from invalid letters results in an error
// Note that only the first error is returned, not all errors
// The invalid '%' is caught, but nothing is said of the invalid '&'
func ExampleNewProtein_errored() {
	s, err := immutable.NewProtein("%" + "ACDEFGHIKLMNPQRSTVWY" + "&")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// %ACDEFGHIKLMNPQRSTVWY&, "%" not in alphabet
}

// Reversing a valid Protein results in no error
func ExampleProtein_Reverse_errorless() {
	s, _ := immutable.NewProtein("ACDEFGHIKLMNPQRSTVWY")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// YWVTSRQPNMLKIHGFEDCA, <nil>
}

// Reversing an invalid Protein results in an error
// Note that only the first error is returned, not all errors
// The invalid '&' is caught, but nothing is said of the invalid '%'
func ExampleProtein_Reverse_errored() {
	s, _ := immutable.NewProtein("%" + "ACDEFGHIKLMNPQRSTVWY" + "&")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// &YWVTSRQPNMLKIHGFEDCA%, "&" not in alphabet
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to immutable.NewProtein()
func ExampleProtein_Alphabet() {
	s, _ := immutable.NewProtein("ACDEFGHIKLMNPQRSTVWY")

	fmt.Println(s.Alphabet())
	// Output:
	// ACDEFGHIKLMNPQRSTVWY
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to immutable.NewProtein()
func ExampleProtein_LetterCount() {
	s, _ := immutable.NewProtein("ACDEFGHIKLMNPQRSTVWY" + "NNNN")

	fmt.Println(s.LetterCount())
	// Output:
	// map[A:1 C:1 D:1 E:1 F:1 G:1 H:1 I:1 K:1 L:1 M:1 N:5 P:1 Q:1 R:1 S:1 T:1 V:1 W:1 Y:1]
}
