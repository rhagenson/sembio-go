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

func TestInitializedDna(t *testing.T) {
	s, _ := NewDna("")
	t.Run("Length is zero", TestLengthIs(s, 0))
	t.Run("Position is empty", TestPositionIs(s, 0, ""))
	t.Run("Range is empty", TestRangeIs(s, 0, 1, ""))
}

func TestDnaHasMethods(t *testing.T) {
	s, _ := NewDna("")

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

func TestDnaCreation(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Dna is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := NewDna(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("Dna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := NewDna(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("Dna has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := NewDna(s)
				onefourth := n * (1 / 4)
				threefourths := n * (3 / 4)
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("Dna has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := NewDna(s)
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

func TestDnaPersistence(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				original, _ := NewDna(s)
				clone := new(Dna)
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
					[]rune(alphabet.Dna.String()),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				original, _ := NewDna(s)
				clone := new(Dna)
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
					[]rune(alphabet.Dna.String()),
				)
				original, _ := NewDna(s)
				clone := new(Dna)
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
					[]rune(alphabet.Dna.String()),
				)
				original, _ := NewDna(s)
				clone := new(Dna)
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
					[]rune(alphabet.Dna.String()),
				)
				original, _ := NewDna(s)
				clone := new(Dna)
				*clone = *original
				original.RevComp()
				return original.seq == clone.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				want, _ := NewDna(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*Dna).Reverse()
				return want.seq == got.(*Dna).seq
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
					[]rune(alphabet.Dna.String()),
				)
				want, _ := NewDna(s)
				rev, _ := want.Complement()
				got, _ := rev.(*Dna).Complement()
				return want.seq == got.(*Dna).seq
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
					[]rune(alphabet.Dna.String()),
				)
				want, _ := NewDna(s)
				rev, _ := want.RevComp()
				got, _ := rev.(*Dna).RevComp()
				return want.seq == got.(*Dna).seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaErrors(t *testing.T) {
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
				if _, err := NewDna(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("Dna creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("Dna should error when using invalid characters, error")
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
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := NewDna(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("Dna should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("Dna Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewDna(s) == NewDna(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				ret := make(chan *Dna)
				go func(s string, out chan *Dna) {
					seq, _ := NewDna(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *Dna) {
					seq, _ := NewDna(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.seq == second.seq
			},
			gen.UIntRange(1, TestableLength),
		),
	)
	properties.Property("seq.(*Dna).Reverse() == seq.(*Dna).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				ret := make(chan *Dna)
				seq, _ := NewDna(s)
				go func(seq *Dna, out chan *Dna) {
					rev, _ := seq.Reverse()
					out <- rev.(*Dna)
				}(seq, ret)
				go func(seq *Dna, out chan *Dna) {
					rev, _ := seq.Reverse()
					out <- rev.(*Dna)
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
					[]rune(alphabet.Dna.String()),
				)
				ret := make(chan *Dna)
				seq, _ := NewDna(s)
				go func(seq *Dna, out chan *Dna) {
					rev, _ := seq.RevComp()
					out <- rev.(*Dna)
				}(seq, ret)
				go func(seq *Dna, out chan *Dna) {
					rev, _ := seq.RevComp()
					out <- rev.(*Dna)
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
					[]rune(alphabet.Dna.String()),
				)
				ret := make(chan *Dna)
				seq, _ := NewDna(s)
				go func(seq *Dna, out chan *Dna) {
					rev, _ := seq.Complement()
					out <- rev.(*Dna)
				}(seq, ret)
				go func(seq *Dna, out chan *Dna) {
					rev, _ := seq.Complement()
					out <- rev.(*Dna)
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
