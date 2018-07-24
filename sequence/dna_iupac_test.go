package sequence

import (
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bigr"
	"bitbucket.org/rhagenson/bigr/alphabet"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedDnaIupac(t *testing.T) {
	s, _ := NewDnaIupac("")
	t.Run("Length is zero", TestLengthIs(s, 0))
	t.Run("Position is empty", TestPositionIs(s, 0, ""))
	t.Run("Range is empty", TestRangeIs(s, 0, 1, ""))
}

func TestDnaIupacHasMethods(t *testing.T) {
	s, _ := NewDnaIupac("")

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

func TestDnaIupacCreation(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("DnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := NewDnaIupac(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("DnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := NewDnaIupac(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("DnaIupac has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := NewDnaIupac(s)
				onefourth := n * (1 / 4)
				threefourths := n * (3 / 4)
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("DnaIupac has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := NewDnaIupac(s)
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

func TestDnaIupacPersistence(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := NewDnaIupac(s)
				clone := new(Sequence)
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
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := NewDnaIupac(s)
				clone := new(Sequence)
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
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := NewDnaIupac(s)
				clone := new(Sequence)
				*clone = *original
				original.Reverse()
				return original.seq == clone.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("Complement does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := NewDnaIupac(s)
				clone := new(Sequence)
				*clone = *original
				original.Complement()
				return original.seq == clone.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("RevComp does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := NewDnaIupac(s)
				clone := new(Sequence)
				*clone = *original
				original.RevComp()
				return original.seq == clone.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				want, _ := NewDnaIupac(s)
				rev, _ := want.Reverse()
				got, _ := rev.Reverse()
				return want.seq == got.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("Complement().Complement() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				want, _ := NewDnaIupac(s)
				rev, _ := want.Complement()
				got, _ := rev.Complement()
				return want.seq == got.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("RevComp().RevComp() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				want, _ := NewDnaIupac(s)
				rev, _ := want.RevComp()
				got, _ := rev.RevComp()
				return want.seq == got.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacErrors(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune("XNQZ"),
				)
				if _, err := NewDnaIupac(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("DnaIupac creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("DnaIupac should error when using invalid characters")
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
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := NewDnaIupac(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("DnaIupac should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("DnaIupac Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewDnaIupac(s) == NewDnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				ret := make(chan *Sequence)
				go func(s string, out chan *Sequence) {
					seq, _ := NewDnaIupac(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *Sequence) {
					seq, _ := NewDnaIupac(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("seq.Reverse() == seq.Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				ret := make(chan *Sequence)
				seq, _ := NewDnaIupac(s)
				go func(seq *Sequence, out chan *Sequence) {
					rev, _ := seq.Reverse()
					out <- rev
				}(seq, ret)
				go func(seq *Sequence, out chan *Sequence) {
					rev, _ := seq.Reverse()
					out <- rev
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("seq.RevComp() == seq.RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				ret := make(chan *Sequence)
				seq, _ := NewDnaIupac(s)
				go func(seq *Sequence, out chan *Sequence) {
					rev, _ := seq.Reverse()
					out <- rev
				}(seq, ret)
				go func(seq *Sequence, out chan *Sequence) {
					rev, _ := seq.Reverse()
					out <- rev
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("seq.Complement() == seq.Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				ret := make(chan *Sequence)
				seq, _ := NewDnaIupac(s)
				go func(seq *Sequence, out chan *Sequence) {
					rev, _ := seq.Reverse()
					out <- rev
				}(seq, ret)
				go func(seq *Sequence, out chan *Sequence) {
					rev, _ := seq.Reverse()
					out <- rev
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
