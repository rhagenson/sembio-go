package immutable_test

import (
	"strings"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/rhagenson/bio-go/bio/alphabet"
	"github.com/rhagenson/bio-go/bio/sequence"
	"github.com/rhagenson/bio-go/bio/sequence/immutable"
	"github.com/rhagenson/bio-go/bio/test"
)

func TestInitializedDnaIupac(t *testing.T) {
	s, _ := immutable.NewDnaIupac("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestDnaIupacHasMethods(t *testing.T) {
	s, _ := immutable.NewDnaIupac("")

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
	t.Run("Has Alphabet method", func(t *testing.T) {
		if a := s.Alphabet(); a == nil {
			t.Error("Alphabet method does not exist")
		}
	})
	t.Run("Has LetterCount method", func(t *testing.T) {
		if c := s.LetterCount(); c == nil {
			t.Error("LetterCount method does not exist")
		}
	})
}

func TestDnaIupacCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("DnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("DnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("DnaIupac has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
				onefourth := n / 4
				threefourths := n * 3 / 4
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("DnaIupac has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
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
	properties.Property("DnaIupac has proper letter count",
		prop.ForAll(
			func(n uint) bool {
				w := n / 4
				s := test.RandomWeightedString(
					test.Seed,
					n,
					map[rune]uint{
						'A': w,
						'T': w,
						'G': w,
						'C': w,
					},
				)
				seq, _ := immutable.NewDnaIupac(s)
				cs := seq.LetterCount()
				for l, c := range cs {
					if c != w {
						t.Errorf("Got %d, want %d for %q.\nSeq: %q\nMap: %v",
							c, w, l, seq, cs)
						return false
					}
				}
				return true
			},
			gen.UIntRange(4, sequence.TestableLength).
				SuchThat(func(u uint) bool { return u%4 == 0 }),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacImmutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.With(immutable.PositionAs(n/2, t))
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
					[]rune(alphabet.DnaIupac.String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.With(immutable.RangeAs(n/4, n*3/4, t))
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
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.Reverse()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.Complement()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				original, _ := immutable.NewDnaIupac(s)
				clone, _ := immutable.NewDnaIupac(s)
				original.RevComp()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				want, _ := immutable.NewDnaIupac(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*immutable.DnaIupac).Reverse()
				return want.String() == got.(*immutable.DnaIupac).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement().Complement() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				want, _ := immutable.NewDnaIupac(s)
				rev, _ := want.Complement()
				got, _ := rev.(*immutable.DnaIupac).Complement()
				return want.String() == got.(*immutable.DnaIupac).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp().RevComp() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				want, _ := immutable.NewDnaIupac(s)
				rev, _ := want.RevComp()
				got, _ := rev.(*immutable.DnaIupac).RevComp()
				return want.String() == got.(*immutable.DnaIupac).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(string(alphabet.TestExcludesLetters([]byte(alphabet.DnaIupac.String())))),
				)
				if _, err := immutable.NewDnaIupac(s); err != nil {
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
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("start > stop errors",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				seq, _ := immutable.NewDnaIupac(s)
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
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("NewDnaIupac(s) == NewDnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				ret := make(chan *immutable.DnaIupac)
				go func(s string, out chan *immutable.DnaIupac) {
					seq, _ := immutable.NewDnaIupac(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *immutable.DnaIupac) {
					seq, _ := immutable.NewDnaIupac(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.Reverse() == seq.Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				ret := make(chan *immutable.DnaIupac)
				seq, _ := immutable.NewDnaIupac(s)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
				}(seq, ret)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.RevComp() == seq.RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				ret := make(chan *immutable.DnaIupac)
				seq, _ := immutable.NewDnaIupac(s)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
				}(seq, ret)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.Complement() == seq.Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(alphabet.DnaIupac.String()),
				)
				ret := make(chan *immutable.DnaIupac)
				seq, _ := immutable.NewDnaIupac(s)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
				}(seq, ret)
				go func(seq *immutable.DnaIupac, out chan *immutable.DnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.DnaIupac)
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
