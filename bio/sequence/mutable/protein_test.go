package mutable_test

import (
	"strings"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/sequence"
	"github.com/rhagenson/bio-go/bio/sequence/mutable"
	"github.com/rhagenson/bio-go/bio/test"
	"github.com/rhagenson/bio-go/bio/utils"
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
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(string(alphabet.TestExcludesLetters([]byte(alphabet.NewProtein().String())))),
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
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
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
					[]rune(alphabet.NewProtein().String()),
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
