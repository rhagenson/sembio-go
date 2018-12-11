package persistent_test

import (
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bio"
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/persistent"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedProtein(t *testing.T) {
	s, _ := persistent.NewProtein("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestProteinHasMethods(t *testing.T) {
	s, _ := persistent.NewProtein("")

	t.Run("Has Reverse method", func(t *testing.T) {
		if _, err := s.Reverse(); err != nil {
			t.Error("Reverse method does not exist")
		}
	})
}

func TestProteinCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Protein is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				seq, _ := persistent.NewProtein(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Protein has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				seq, _ := persistent.NewProtein(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Protein has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				seq, _ := persistent.NewProtein(s)
				onefourth := n * (1 / 4)
				threefourths := n * (3 / 4)
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Protein has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				seq, _ := persistent.NewProtein(s)
				onefourth := n * (1 / 4)
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

func TestProteinPersistence(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				t := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				original, _ := persistent.NewProtein(s)
				clone := new(persistent.Protein)
				*clone = *original
				original.With(persistent.PositionAs(n*(1/2), t))
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("WithRange does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				t := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				original, _ := persistent.NewProtein(s)
				clone := new(persistent.Protein)
				*clone = *original
				original.With(persistent.RangeAs(n*(1/4), n*(3/4), t))
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Reverse does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				original, _ := persistent.NewProtein(s)
				clone := new(persistent.Protein)
				*clone = *original
				original.Reverse()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				want, _ := persistent.NewProtein(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*persistent.Protein).Reverse()
				return want.String() == got.(*persistent.Protein).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune("XNQZ"),
				)
				if _, err := persistent.NewProtein(s); err != nil {
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
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				seq, _ := persistent.NewProtein(s)
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
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("persistent.NewProtein(s) == persistent.NewProtein(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				ret := make(chan *persistent.Protein)
				go func(s string, out chan *persistent.Protein) {
					seq, _ := persistent.NewProtein(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *persistent.Protein) {
					seq, _ := persistent.NewProtein(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.(*persistent.Protein).Reverse() == seq.(*persistent.Protein).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Protein.String()),
				)
				ret := make(chan *persistent.Protein)
				seq, _ := persistent.NewProtein(s)
				go func(seq *persistent.Protein, out chan *persistent.Protein) {
					rev, _ := seq.Reverse()
					out <- rev.(*persistent.Protein)
				}(seq, ret)
				go func(seq *persistent.Protein, out chan *persistent.Protein) {
					rev, _ := seq.Reverse()
					out <- rev.(*persistent.Protein)
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
