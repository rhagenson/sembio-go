package sequence

import (
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bio"
	"bitbucket.org/rhagenson/bio/alphabet"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedProteinGapped(t *testing.T) {
	s, _ := NewProteinGapped("")
	t.Run("Length is zero", TestLengthIs(s, 0))
	t.Run("Position is empty", TestPositionIs(s, 0, ""))
	t.Run("Range is empty", TestRangeIs(s, 0, 1, ""))
}

func TestProteinGappedHasMethods(t *testing.T) {
	s, _ := NewProteinGapped("")

	t.Run("Has Reverse method", func(t *testing.T) {
		if _, err := s.Reverse(); err != nil {
			t.Error("Reverse method does not exist")
		}
	})
}

func TestProteinGappedCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("ProteinGapped is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				seq, _ := NewProteinGapped(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("ProteinGapped has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				seq, _ := NewProteinGapped(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("ProteinGapped has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				seq, _ := NewProteinGapped(s)
				onefourth := n * (1 / 4)
				threefourths := n * (3 / 4)
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("ProteinGapped has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				seq, _ := NewProteinGapped(s)
				onefourth := n * (1 / 4)
				threefourth := n * (3 / 4)
				gotoneforth, _ := seq.Position(onefourth)
				wantoneforth := string(s[onefourth])
				gotthreeforth, _ := seq.Position(threefourth)
				wantthreeforth := string(s[threefourth])
				return gotoneforth == wantoneforth && gotthreeforth == wantthreeforth
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinGappedPersistence(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				t := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				original, _ := NewProteinGapped(s)
				clone := new(ProteinGapped)
				*clone = *original
				original.With(PositionAs(n*(1/2), t))
				return original.seq == clone.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("WithRange does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				t := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				original, _ := NewProteinGapped(s)
				clone := new(ProteinGapped)
				*clone = *original
				original.With(RangeAs(n*(1/4), n*(3/4), t))
				return original.seq == clone.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("Reverse does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				original, _ := NewProteinGapped(s)
				clone := new(ProteinGapped)
				*clone = *original
				original.Reverse()
				return original.seq == clone.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinGappedMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				want, _ := NewProteinGapped(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*ProteinGapped).Reverse()
				return want.seq == got.(*ProteinGapped).seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinGappedErrors(t *testing.T) {
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
				if _, err := NewProteinGapped(s); err != nil {
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
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("start > stop errors",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				seq, _ := NewProteinGapped(s)
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
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestProteinGappedParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("NewProteinGapped(s) == NewProteinGapped(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				ret := make(chan *ProteinGapped)
				go func(s string, out chan *ProteinGapped) {
					seq, _ := NewProteinGapped(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *ProteinGapped) {
					seq, _ := NewProteinGapped(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("seq.(*ProteinGapped).Reverse() == seq.(*ProteinGapped).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.ProteinGapped.String()),
				)
				ret := make(chan *ProteinGapped)
				seq, _ := NewProteinGapped(s)
				go func(seq *ProteinGapped, out chan *ProteinGapped) {
					rev, _ := seq.Reverse()
					out <- rev.(*ProteinGapped)
				}(seq, ret)
				go func(seq *ProteinGapped, out chan *ProteinGapped) {
					rev, _ := seq.Reverse()
					out <- rev.(*ProteinGapped)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}
