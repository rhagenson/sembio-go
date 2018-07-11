package sequence

import (
	"reflect"
	"testing"

	"bitbucket.org/rhagenson/bigr"
	"bitbucket.org/rhagenson/bigr/alphabet"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

var _ DnaStrictSequence = new(SimpleDna)

func TestInitializedSimpleDNA(t *testing.T) {
	dna := new(SimpleDna)

	if dna.Alphabet() != new(alphabet.DNAStrict) {
		t.Errorf("Want: %t, Got: %t", new(alphabet.DNAStrict), dna.Alphabet())
	}
	if dna.Length() != 0 {
		t.Errorf("Want: %d, Got: %d", 0, dna.Length())
	}
	// TODO: Write test for runtime panic on dna.Postion() and dna.Range()
}

func TestSimpleDnaCreation(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("SimpleDna is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaStrictLetters),
				)
				dna := NewSimpleDna(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, seqLen),
		),
	)
	properties.Property("SimpleDna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaStrictLetters),
				)
				dna := NewSimpleDna(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestSimpleDnaPersistence(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaStrictLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaStrictLetters),
				)
				original := NewSimpleDna(s)
				clone := new(SimpleDna)
				*clone = *original
				mut := original.WithPosition(n*(1/2), t)
				return reflect.DeepEqual(original, clone) &&
					!reflect.DeepEqual(original, mut)
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("WithRange does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaStrictLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaStrictLetters),
				)
				original := NewSimpleDna(s)
				clone := new(SimpleDna)
				*clone = *original
				mut := original.WithRange(n*(1/4), n*(3/4), t)
				return reflect.DeepEqual(original, clone) &&
					!reflect.DeepEqual(original, mut)
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}
