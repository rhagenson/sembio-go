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
	_ sequence.Interface = NewDna("")
)

func TestInitializedDna(t *testing.T) {
	s := NewDna("")

	t.Run("DNA alphabet",
		sequence.TestAlphabetIs(s.Alphabet(), new(simple.Dna)),
	)
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
}

func TestDnaHasMethods(t *testing.T) {
	s := NewDna("")

	t.Run("Has Reverse method", bigr.TestForMethodNamed(s, "Reverse"))
	t.Run("Has Complement method", bigr.TestForMethodNamed(s, "Complement"))
	t.Run("Has RevComp method", bigr.TestForMethodNamed(s, "RevComp"))
	t.Run("Has Alphabet method", bigr.TestForMethodNamed(s, "Alphabet"))
}

func TestDnaMethodReturnType(t *testing.T) {
	s := NewDna("")

	t.Run("Reverse returns *Dna",
		bigr.TestMethodReturnsSelfType(s, "Reverse", nil),
	)
	t.Run("Complement returns *Dna",
		bigr.TestMethodReturnsSelfType(s, "Complement", nil),
	)
	t.Run("RevComp returns *Dna",
		bigr.TestMethodReturnsSelfType(s, "RevComp", nil),
	)
	t.Run("Alphabet returns *simple.Dna",
		bigr.TestMethodReturnsType(s, new(simple.Dna), "Alphabet", nil),
	)
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
					[]rune(simple.DnaLetters),
				)
				dna := NewDna(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				dna := NewDna(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				dna := NewDna(s)
				onefourth := n * (1 / 4)
				threefourths := n * (3 / 4)
				got := dna.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				dna := NewDna(s)
				onefourth := n * (1 / 4)
				threefourth := n * (3 / 4)
				gotoneforth := dna.Position(onefourth)
				wantoneforth := string(s[onefourth])
				gotthreeforth := dna.Position(threefourth)
				wantthreeforth := string(s[threefourth])
				return gotoneforth == wantoneforth && gotthreeforth == wantthreeforth
			},
			gen.UIntRange(1, sequence.TestableLength),
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
					[]rune(simple.DnaLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
				*clone = *original
				_ = original.WithPosition(n*(1/2), t)
				return reflect.DeepEqual(original, clone)
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
					[]rune(simple.DnaLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
				*clone = *original
				_ = original.WithRange(n*(1/4), n*(3/4), t)
				return reflect.DeepEqual(original, clone)
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
					[]rune(simple.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
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
					[]rune(simple.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
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
					[]rune(simple.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
				*clone = *original
				_ = original.RevComp()
				return reflect.DeepEqual(original, clone)
			},
			gen.UIntRange(1, sequence.TestableLength),
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
					[]rune(simple.DnaLetters),
				)
				original := NewDna(s)
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
					[]rune(simple.DnaLetters),
				)
				original := NewDna(s)
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
					[]rune(simple.DnaLetters),
				)
				original := NewDna(s)
				return reflect.DeepEqual(original, original.RevComp().RevComp())
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaAccumulatesErrors(t *testing.T) {
	var _ ErrorAccumulator = new(Dna)
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
				seq := NewDna(s)
				for _, err := range seq.Errors() {
					if err == nil {
						t.Errorf("Dna should accumulate an err using non-standard chars")
						return false
					}
					if !strings.Contains(err.Error(), "invalid character(s)") {
						t.Errorf("Dna creation error should mention invalid character(s)")
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
					[]rune(simple.DnaLetters),
				)
				seq := NewDna(s)
				seq.Range(n, 0)
				for _, err := range seq.Errors() {
					if err == nil {
						t.Errorf("Dna should accumulate an err during Range() when start > stop")
						return false
					}
					if !strings.Contains(err.Error(), "impossible range") {
						t.Errorf("Dna Range error should mention impossible range")
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

func TestDnaParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewDna(s) == NewDna(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				left := make(chan *Dna)
				right := make(chan *Dna)
				go func(s string, out chan *Dna) {
					out <- NewDna(s)
				}(s, left)
				go func(s string, out chan *Dna) {
					out <- NewDna(s)
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewDna(s).Reverse() == NewDna(s).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				left := make(chan *Dna)
				right := make(chan *Dna)
				go func(s string, out chan *Dna) {
					out <- NewDna(s).Reverse()
				}(s, left)
				go func(s string, out chan *Dna) {
					out <- NewDna(s).Reverse()
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewDna(s).RevComp() == NewDna(s).RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				left := make(chan *Dna)
				right := make(chan *Dna)
				go func(s string, out chan *Dna) {
					out <- NewDna(s).RevComp()
				}(s, left)
				go func(s string, out chan *Dna) {
					out <- NewDna(s).RevComp()
				}(s, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("NewDna(s).Complement() == NewDna(s).Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(simple.DnaLetters),
				)
				left := make(chan *Dna)
				right := make(chan *Dna)
				seq := NewDna(s)
				go func(seq *Dna, out chan *Dna) {
					out <- seq.Complement()
				}(seq, left)
				go func(seq *Dna, out chan *Dna) {
					out <- seq.Complement()
				}(seq, right)
				return reflect.DeepEqual(<-left, <-right)
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}
