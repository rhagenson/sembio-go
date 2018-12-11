package persistent

import (
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bio"
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/sequence"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedRnaIupac(t *testing.T) {
	s, _ := NewRnaIupac("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestRnaIupacHasMethods(t *testing.T) {
	s, _ := NewRnaIupac("")

	t.Run("Has Reverse method", func(t *testing.T) {
		if _, err := s.Reverse(); err != nil {
			t.Error("Reverse method does not exist")
		}
	})
	t.Run("Has Complement method", func(t *testing.T) {
		if _, err := s.Complement(); err != nil {
			t.Error("Complement method does not exist")
		}
	})
	t.Run("Has RevComp method", func(t *testing.T) {
		if _, err := s.RevComp(); err != nil {
			t.Error("RevComp method does not exist")
		}
	})
}

func TestRnaIupacCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("RnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				seq, _ := NewRnaIupac(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				seq, _ := NewRnaIupac(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				seq, _ := NewRnaIupac(s)
				onefourth := n * (1 / 4)
				threefourths := n * (3 / 4)
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				seq, _ := NewRnaIupac(s)
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

func TestRnaIupacPersistence(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				t := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				original.With(PositionAs(n*(1/2), t))
				return original.seq == clone.seq
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
					[]rune(alphabet.RnaIupac.String()),
				)
				t := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				original.With(RangeAs(n*(1/4), n*(3/4), t))
				return original.seq == clone.seq
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
					[]rune(alphabet.RnaIupac.String()),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				original.Reverse()
				return original.seq == clone.seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				original.Complement()
				return original.seq == clone.seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				original.RevComp()
				return original.seq == clone.seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				want, _ := NewRnaIupac(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*RnaIupac).Reverse()
				return want.seq == got.(*RnaIupac).seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement().Complement() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				want, _ := NewRnaIupac(s)
				rev, _ := want.Complement()
				got, _ := rev.(*RnaIupac).Complement()
				return want.seq == got.(*RnaIupac).seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp().RevComp() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				want, _ := NewRnaIupac(s)
				rev, _ := want.RevComp()
				got, _ := rev.(*RnaIupac).RevComp()
				return want.seq == got.(*RnaIupac).seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacErrors(t *testing.T) {
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
				if _, err := NewRnaIupac(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("RnaIupac creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("RnaIupac should error when using invalid characters")
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
					[]rune(alphabet.RnaIupac.String()),
				)
				seq, _ := NewRnaIupac(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("RnaIupac should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("RnaIupac Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("NewRnaIupac(s) == NewRnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				ret := make(chan *RnaIupac)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.Reverse() == seq.Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				ret := make(chan *RnaIupac)
				seq, _ := NewRnaIupac(s)
				go func(seq *RnaIupac, out chan *RnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*RnaIupac)
				}(seq, ret)
				go func(seq *RnaIupac, out chan *RnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*RnaIupac)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.RevComp() == seq.RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				ret := make(chan *RnaIupac)
				seq, _ := NewRnaIupac(s)
				go func(seq *RnaIupac, out chan *RnaIupac) {
					rev, _ := seq.RevComp()
					out <- rev.(*RnaIupac)
				}(seq, ret)
				go func(seq *RnaIupac, out chan *RnaIupac) {
					rev, _ := seq.RevComp()
					out <- rev.(*RnaIupac)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.Complement() == seq.Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.RnaIupac.String()),
				)
				ret := make(chan *RnaIupac)
				seq, _ := NewRnaIupac(s)
				go func(seq *RnaIupac, out chan *RnaIupac) {
					rev, _ := seq.Complement()
					out <- rev.(*RnaIupac)
				}(seq, ret)
				go func(seq *RnaIupac, out chan *RnaIupac) {
					rev, _ := seq.Complement()
					out <- rev.(*RnaIupac)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}
