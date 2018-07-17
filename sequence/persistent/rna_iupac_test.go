package persistent

import (
	"reflect"
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bigr"
	"bitbucket.org/rhagenson/bigr/alphabet/simple"
	"bitbucket.org/rhagenson/bigr/sequence"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

var (
	_ sequence.Interface = new(RnaIupac)
)

func TestInitializedRnaIupac(t *testing.T) {
	s, _ := NewRnaIupac("")

	t.Run("IUPAC RNA alphabet",
		sequence.TestAlphabetIs(s.Alphabet(), new(simple.RnaIupac)),
	)
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestRnaIupacHasMethods(t *testing.T) {
	s, _ := NewRnaIupac("")

	t.Run("Has Reverse method", bigr.TestForMethodNamed(s, "Reverse"))
	t.Run("Has Complement method", bigr.TestForMethodNamed(s, "Complement"))
	t.Run("Has RevComp method", bigr.TestForMethodNamed(s, "RevComp"))
	t.Run("Has Alphabet method", bigr.TestForMethodNamed(s, "Alphabet"))
}

func TestRnaIupacMethodsReturnTypes(t *testing.T) {
	s, _ := NewRnaIupac("")

	t.Run("Reverse returns *RnaIupac",
		bigr.TestMethodReturnsSelfType(s, "Reverse", nil),
	)
	t.Run("Complement returns *RnaIupac",
		bigr.TestMethodReturnsSelfType(s, "Complement", nil),
	)
	t.Run("RevComp returns *RnaIupac",
		bigr.TestMethodReturnsSelfType(s, "RevComp", nil),
	)
	t.Run("Alphabet returns *simple.RnaIupac",
		bigr.TestMethodReturnsType(s, new(simple.RnaIupac), "Alphabet", nil),
	)
}

func TestRnaIupacCreation(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("RnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				rna, _ := NewRnaIupac(s)
				return rna.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				rna, _ := NewRnaIupac(s)
				got, _ := rna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				rna, _ := NewRnaIupac(s)
				onefourth := n * (1 / 4)
				threefourths := n * (3 / 4)
				got, _ := rna.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RnaIupac has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				rna, _ := NewRnaIupac(s)
				onefourth := n * (1 / 4)
				threefourth := n * (3 / 4)
				gotoneforth, _ := rna.Position(onefourth)
				wantoneforth := string(s[onefourth])
				gotthreeforth, _ := rna.Position(threefourth)
				wantthreeforth := string(s[threefourth])
				return gotoneforth == wantoneforth && gotthreeforth == wantthreeforth
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacPersistence(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				mut, _ := original.WithPosition(n*(1/2), t)
				return reflect.DeepEqual(original, clone) &&
					!reflect.DeepEqual(original, mut)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("WithRange does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				mut, _ := original.WithRange(n*(1/4), n*(3/4), t)
				return reflect.DeepEqual(original, clone) &&
					!reflect.DeepEqual(original, mut)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Reverse does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				original.Reverse()
				return reflect.DeepEqual(original, clone)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				original.Complement()
				return reflect.DeepEqual(original, clone)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				original, _ := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				original.RevComp()
				return reflect.DeepEqual(original, clone)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				want, _ := NewRnaIupac(s)
				rev, _ := want.Reverse()
				got, _ := rev.Reverse()
				return reflect.DeepEqual(want, got)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement().Complement() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				want, _ := NewRnaIupac(s)
				rev, _ := want.Complement()
				got, _ := rev.Complement()
				return reflect.DeepEqual(want, got)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp().RevComp() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				want, _ := NewRnaIupac(s)
				rev, _ := want.RevComp()
				got, _ := rev.RevComp()
				return reflect.DeepEqual(want, got)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacAccumulatesErrors(t *testing.T) {
	var _ ErrorAccumulator = new(RnaIupac)
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
				_, err := NewRnaIupac(s)
				if err == nil {
					t.Errorf("RnaIupac should accumulate an err using non-standard chars")
					return false
				}
				if !strings.Contains(err.Error(), "invalid character(s)") {
					t.Errorf("RnaIupac creation error should mention invalid character(s)")
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
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
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
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewRnaIupac(s) == NewRnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				left := make(chan *RnaIupac)
				right := make(chan *RnaIupac)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					out <- seq
				}(s, left)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					out <- seq
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewRnaIupac(s).Reverse() == NewRnaIupac(s).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				left := make(chan *RnaIupac)
				right := make(chan *RnaIupac)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					rev, _ := seq.Reverse()
					out <- rev
				}(s, left)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					rev, _ := seq.Reverse()
					out <- rev
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewRnaIupac(s).RevComp() == NewRnaIupac(s).RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				left := make(chan *RnaIupac)
				right := make(chan *RnaIupac)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					revcomp, _ := seq.RevComp()
					out <- revcomp
				}(s, left)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					revcomp, _ := seq.RevComp()
					out <- revcomp
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewRnaIupac(s).Complement() == NewRnaIupac(s).Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaIupacLetters),
				)
				left := make(chan *RnaIupac)
				right := make(chan *RnaIupac)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					comp, _ := seq.Complement()
					out <- comp
				}(s, left)
				go func(s string, out chan *RnaIupac) {
					seq, _ := NewRnaIupac(s)
					comp, _ := seq.Complement()
					out <- comp
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}
