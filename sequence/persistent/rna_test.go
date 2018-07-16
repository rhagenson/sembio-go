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
	_ sequence.Interface = NewRna("")
)

func TestInitializedRna(t *testing.T) {
	s := NewRna("")

	t.Run("RNA alphabet",
		sequence.TestAlphabetIs(s.Alphabet(), new(simple.Rna)),
	)
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestRnaHasMethods(t *testing.T) {
	s := NewRna("")

	t.Run("Has Reverse method", bigr.TestForMethodNamed(s, "Reverse"))
	t.Run("Has Complement method", bigr.TestForMethodNamed(s, "Complement"))
	t.Run("Has RevComp method", bigr.TestForMethodNamed(s, "RevComp"))
	t.Run("Has Alphabet method", bigr.TestForMethodNamed(s, "Alphabet"))
}

func TestRnaMethodsReturnTypes(t *testing.T) {
	s := NewRna("")

	t.Run("Reverse returns *Rna",
		bigr.TestMethodReturnsSelfType(s, "Reverse", nil),
	)
	t.Run("Complement returns *Rna",
		bigr.TestMethodReturnsSelfType(s, "Complement", nil),
	)
	t.Run("RevComp returns *Rna",
		bigr.TestMethodReturnsSelfType(s, "RevComp", nil),
	)
	t.Run("Alphabet returns *simple.Rna",
		bigr.TestMethodReturnsType(s, new(simple.Rna), "Alphabet", nil),
	)
}

func TestRnaCreation(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Rna is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				dna := NewRna(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Rna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				dna := NewRna(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaPersistence(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				original := NewRna(s)
				clone := NewRna("")
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
					[]rune(simple.RnaLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				original := NewRna(s)
				clone := NewRna("")
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
					[]rune(simple.RnaLetters),
				)
				original := NewRna(s)
				clone := new(Rna)
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
					[]rune(simple.RnaLetters),
				)
				original := NewRna(s)
				clone := new(Rna)
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
					[]rune(simple.RnaLetters),
				)
				original := NewRna(s)
				clone := new(Rna)
				*clone = *original
				_ = original.RevComp()
				return reflect.DeepEqual(original, clone)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				original := NewRna(s)
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
					[]rune(simple.RnaLetters),
				)
				original := NewRna(s)
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
					[]rune(simple.RnaLetters),
				)
				original := NewRna(s)
				return reflect.DeepEqual(original, original.RevComp().RevComp())
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestRnaAccumulatesErrors(t *testing.T) {
	var _ ErrorAccumulator = NewRna("")
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
				seq := NewRna(s)
				for _, err := range seq.Errors() {
					if err == nil {
						t.Errorf("Rna should accumulate an err using non-standard chars")
						return false
					}
					if !strings.Contains(err.Error(), "invalid character(s)") {
						t.Errorf("Rna creation error should mention invalid character(s)")
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
					[]rune(simple.RnaLetters),
				)
				seq := NewRna(s)
				seq.Range(n, 0)
				for _, err := range seq.Errors() {
					if err == nil {
						t.Errorf("Rna should accumulate an err during Range() when start > stop")
						return false
					}
					if !strings.Contains(err.Error(), "impossible range") {
						t.Errorf("Rna Range error should mention impossible range")
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

func TestRnaParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewRna(s) == NewRna(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				left := make(chan *Rna)
				right := make(chan *Rna)
				go func(s string, out chan *Rna) {
					out <- NewRna(s)
				}(s, left)
				go func(s string, out chan *Rna) {
					out <- NewRna(s)
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewRna(s).Reverse() == NewRna(s).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				left := make(chan *Rna)
				right := make(chan *Rna)
				go func(s string, out chan *Rna) {
					out <- NewRna(s).Reverse()
				}(s, left)
				go func(s string, out chan *Rna) {
					out <- NewRna(s).Reverse()
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewRna(s).RevComp() == NewRna(s).RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				left := make(chan *Rna)
				right := make(chan *Rna)
				go func(s string, out chan *Rna) {
					out <- NewRna(s).RevComp()
				}(s, left)
				go func(s string, out chan *Rna) {
					out <- NewRna(s).RevComp()
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewRna(s).Complement() == NewRna(s).Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.RnaLetters),
				)
				left := make(chan *Rna)
				right := make(chan *Rna)
				seq := NewRna(s)
				go func(seq *Rna, out chan *Rna) {
					out <- seq.Complement()
				}(seq, left)
				go func(seq *Rna, out chan *Rna) {
					out <- seq.Complement()
				}(seq, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}
