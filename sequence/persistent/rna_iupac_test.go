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
	_ sequence.Interface = NewRnaIupac("")
)

func TestInitializedRnaIupac(t *testing.T) {
	s := NewRnaIupac("")

	t.Run("IUPAC RNA alphabet",
		sequence.TestAlphabetIs(s.Alphabet(), new(simple.RnaIupac)),
	)
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestRnaIupacHasMethods(t *testing.T) {
	s := NewRnaIupac("")

	t.Run("Has Reverse method", bigr.TestForMethodNamed(s, "Reverse"))
	t.Run("Has Complement method", bigr.TestForMethodNamed(s, "Complement"))
	t.Run("Has RevComp method", bigr.TestForMethodNamed(s, "RevComp"))
	t.Run("Has Alphabet method", bigr.TestForMethodNamed(s, "Alphabet"))
}

func TestRnaIupacMethodsReturnTypes(t *testing.T) {
	s := NewRnaIupac("")

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
				dna := NewRnaIupac(s)
				return dna.Length() == n
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
				dna := NewRnaIupac(s)
				got := dna.Range(0, n)
				return got == s
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
				original := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				mut := original.WithPosition(n*(1/2), t)
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
				original := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				mut := original.WithRange(n*(1/4), n*(3/4), t)
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
				original := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				_ = original.Reverse()
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
				original := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				_ = original.Complement()
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
				original := NewRnaIupac(s)
				clone := new(RnaIupac)
				*clone = *original
				_ = original.RevComp()
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
				original := NewRnaIupac(s)
				return reflect.DeepEqual(original, original.Reverse().Reverse())
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
				original := NewRnaIupac(s)
				return reflect.DeepEqual(original, original.Complement().Complement())
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
				original := NewRnaIupac(s)
				return reflect.DeepEqual(original, original.RevComp().RevComp())
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacAccumulatesErrors(t *testing.T) {
	var _ ErrorAccumulator = NewRnaIupac("")
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
				seq := NewRnaIupac(s)
				for _, err := range seq.errs {
					if err == nil {
						t.Errorf("RnaIupac should accumulate an err using non-standard chars")
						return false
					}
					if !strings.Contains(err.Error(), "invalid character(s)") {
						t.Errorf("RnaIupac creation error should mention invalid character(s)")
						return false
					}
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
				seq := NewRnaIupac(s)
				seq.Range(n, 0)
				for _, err := range seq.errs {
					if err == nil {
						t.Errorf("RnaIupac should accumulate an err during Range() when start > stop")
						return false
					}
					if !strings.Contains(err.Error(), "impossible range") {
						t.Errorf("RnaIupac Range error should mention impossible range")
						return false
					}
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
					out <- NewRnaIupac(s)
				}(s, left)
				go func(s string, out chan *RnaIupac) {
					out <- NewRnaIupac(s)
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
					out <- NewRnaIupac(s).Reverse()
				}(s, left)
				go func(s string, out chan *RnaIupac) {
					out <- NewRnaIupac(s).Reverse()
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
					out <- NewRnaIupac(s).RevComp()
				}(s, left)
				go func(s string, out chan *RnaIupac) {
					out <- NewRnaIupac(s).RevComp()
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
				seq := NewRnaIupac(s)
				go func(seq *RnaIupac, out chan *RnaIupac) {
					out <- seq.Complement()
				}(seq, left)
				go func(seq *RnaIupac, out chan *RnaIupac) {
					out <- seq.Complement()
				}(seq, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}
