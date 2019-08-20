package immutable_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/sequence"
	"github.com/bio-ext/bio-go/bio/sequence/immutable"
	"github.com/bio-ext/bio-go/bio/test"
)

func TestInitializedRnaIupac(t *testing.T) {
	s, _ := immutable.NewRnaIupac("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestRnaIupacHasMethods(t *testing.T) {
	s, _ := immutable.NewRnaIupac("")

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
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("RnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := immutable.NewRnaIupac(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := immutable.NewRnaIupac(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := immutable.NewRnaIupac(s)
				onefourth := n / 4
				threefourths := n * 3 / 4
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := immutable.NewRnaIupac(s)
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

func TestRnaIupacImmutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := immutable.NewRnaIupac(s)
				clone, _ := immutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := immutable.NewRnaIupac(s)
				clone, _ := immutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := immutable.NewRnaIupac(s)
				clone, _ := immutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := immutable.NewRnaIupac(s)
				clone, _ := immutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				original, _ := immutable.NewRnaIupac(s)
				clone, _ := immutable.NewRnaIupac(s)
				original.RevComp()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				want, _ := immutable.NewRnaIupac(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*immutable.RnaIupac).Reverse()
				return want.String() == got.(*immutable.RnaIupac).String()
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				want, _ := immutable.NewRnaIupac(s)
				rev, _ := want.Complement()
				got, _ := rev.(*immutable.RnaIupac).Complement()
				return want.String() == got.(*immutable.RnaIupac).String()
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				orig, _ := immutable.NewRnaIupac(s)
				want := orig.String()
				rev, _ := orig.RevComp()
				drev, _ := rev.(*immutable.RnaIupac).RevComp()
				got, _ := drev.Range(0, drev.Length())
				fmt.Printf("want: %s\ngot: %s\n", want, got)
				return want == got
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(string(alphabet.TestExcludesSingleLetters([]byte(hashmap.NewRnaIupac().String())))),
				)
				if _, err := immutable.NewRnaIupac(s); err != nil {
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
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				seq, _ := immutable.NewRnaIupac(s)
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
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("immutable.NewRnaIupac(s) == immutable.NewRnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewRnaIupac().String()),
				)
				ret := make(chan *immutable.RnaIupac)
				go func(s string, out chan *immutable.RnaIupac) {
					seq, _ := immutable.NewRnaIupac(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *immutable.RnaIupac) {
					seq, _ := immutable.NewRnaIupac(s)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				ret := make(chan *immutable.RnaIupac)
				seq, _ := immutable.NewRnaIupac(s)
				go func(seq *immutable.RnaIupac, out chan *immutable.RnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.RnaIupac)
				}(seq, ret)
				go func(seq *immutable.RnaIupac, out chan *immutable.RnaIupac) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.RnaIupac)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				ret := make(chan *immutable.RnaIupac)
				seq, _ := immutable.NewRnaIupac(s)
				go func(seq *immutable.RnaIupac, out chan *immutable.RnaIupac) {
					rev, _ := seq.RevComp()
					out <- rev.(*immutable.RnaIupac)
				}(seq, ret)
				go func(seq *immutable.RnaIupac, out chan *immutable.RnaIupac) {
					rev, _ := seq.RevComp()
					out <- rev.(*immutable.RnaIupac)
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
					[]rune(hashmap.NewRnaIupac().String()),
				)
				ret := make(chan *immutable.RnaIupac)
				seq, _ := immutable.NewRnaIupac(s)
				go func(seq *immutable.RnaIupac, out chan *immutable.RnaIupac) {
					rev, _ := seq.Complement()
					out <- rev.(*immutable.RnaIupac)
				}(seq, ret)
				go func(seq *immutable.RnaIupac, out chan *immutable.RnaIupac) {
					rev, _ := seq.Complement()
					out <- rev.(*immutable.RnaIupac)
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
