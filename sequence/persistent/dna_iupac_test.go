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
	_ sequence.Interface = NewDnaIupac("")
)

func TestInitializedDnaIupac(t *testing.T) {
	s := NewDnaIupac("")

	t.Run("IUPAC DNA alphabet",
		sequence.TestAlphabetIs(s.Alphabet(), new(simple.DnaIupac)),
	)
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestDnaIupacHasMethods(t *testing.T) {
	s := NewDnaIupac("")

	t.Run("Has Reverse method", bigr.TestForMethodNamed(s, "Reverse"))
	t.Run("Has Complement method", bigr.TestForMethodNamed(s, "Complement"))
	t.Run("Has RevComp method", bigr.TestForMethodNamed(s, "RevComp"))
	t.Run("Has Alphabet method", bigr.TestForMethodNamed(s, "Alphabet"))
}

func TestDnaIupacMethodReturnType(t *testing.T) {
	s := NewDnaIupac("")

	t.Run("Reverse returns *DnaIupac",
		bigr.TestMethodReturnsSelfType(s, "Reverse", nil),
	)
	t.Run("Complement returns *DnaIupac",
		bigr.TestMethodReturnsSelfType(s, "Complement", nil),
	)
	t.Run("RevComp returns *DnaIupac",
		bigr.TestMethodReturnsSelfType(s, "RevComp", nil),
	)
	t.Run("Alphabet returns *simple.DnaIupac",
		bigr.TestMethodReturnsType(s, new(simple.DnaIupac), "Alphabet", nil),
	)
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
					[]rune(simple.DnaIupacLetters),
				)
				dna := NewDnaIupac(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("DnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaIupacLetters),
				)
				dna := NewDnaIupac(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
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
					[]rune(simple.DnaIupacLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
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
					[]rune(simple.DnaIupacLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
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
					[]rune(simple.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
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
					[]rune(simple.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
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
					[]rune(simple.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
				*clone = *original
				_ = original.RevComp()
				return reflect.DeepEqual(original, clone)
			},
			gen.UIntRange(1, sequence.TestableLength),
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
					[]rune(simple.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
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
					[]rune(simple.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
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
					[]rune(simple.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				return reflect.DeepEqual(original, original.RevComp().RevComp())
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacAccumulatesErrors(t *testing.T) {
	var _ ErrorAccumulator = NewDnaIupac("")
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Invalid input errors",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune("XNQZ"),
				)
				seq := NewDnaIupac(s)
				for _, err := range seq.errs {
					if err == nil {
						t.Errorf("DnaIupac should accumulate an err using non-standard chars")
						return false
					}
					if !strings.Contains(err.Error(), "invalid character(s)") {
						t.Errorf("DnaIupac creation error should mention invalid character(s)")
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
					[]rune(simple.DnaIupacLetters),
				)
				seq := NewDnaIupac(s)
				seq.Range(n, 0)
				for _, err := range seq.errs {
					if err == nil {
						t.Errorf("DnaIupac should accumulate an err during Range() when start > stop")
						return false
					}
					if !strings.Contains(err.Error(), "impossible range") {
						t.Errorf("DnaIupac Range error should mention impossible range")
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

func TestDnaIupacParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewDnaIupac(s) == NewDnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaIupacLetters),
				)
				left := make(chan *DnaIupac)
				right := make(chan *DnaIupac)
				go func(s string, out chan *DnaIupac) {
					out <- NewDnaIupac(s)
				}(s, left)
				go func(s string, out chan *DnaIupac) {
					out <- NewDnaIupac(s)
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewDnaIupac(s).Reverse() == NewDnaIupac(s).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaIupacLetters),
				)
				left := make(chan *DnaIupac)
				right := make(chan *DnaIupac)
				go func(s string, out chan *DnaIupac) {
					out <- NewDnaIupac(s).Reverse()
				}(s, left)
				go func(s string, out chan *DnaIupac) {
					out <- NewDnaIupac(s).Reverse()
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewDnaIupac(s).RevComp() == NewDnaIupac(s).RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaIupacLetters),
				)
				left := make(chan *DnaIupac)
				right := make(chan *DnaIupac)
				go func(s string, out chan *DnaIupac) {
					out <- NewDnaIupac(s).RevComp()
				}(s, left)
				go func(s string, out chan *DnaIupac) {
					out <- NewDnaIupac(s).RevComp()
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewDnaIupac(s).Complement() == NewDnaIupac(s).Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaIupacLetters),
				)
				left := make(chan *DnaIupac)
				right := make(chan *DnaIupac)
				seq := NewDnaIupac(s)
				go func(seq *DnaIupac, out chan *DnaIupac) {
					out <- seq.Complement()
				}(seq, left)
				go func(seq *DnaIupac, out chan *DnaIupac) {
					out <- seq.Complement()
				}(seq, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}
