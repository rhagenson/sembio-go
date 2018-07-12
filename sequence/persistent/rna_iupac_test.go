package persistent

import (
	"reflect"
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bigr"
	"bitbucket.org/rhagenson/bigr/alphabet"
	"bitbucket.org/rhagenson/bigr/helpers"
	"bitbucket.org/rhagenson/bigr/sequence"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

var (
	_ sequence.Interface = new(RnaIupac)
)

func TestInitializedRnaIupac(t *testing.T) {
	rna := new(RnaIupac)

	if rna.Alphabet() != new(alphabet.RnaIupac) {
		t.Errorf("Want: %t, Got: %t", new(alphabet.RnaIupac), rna.Alphabet())
	}
	if rna.Length() != 0 {
		t.Errorf("Want: %d, Got: %d", 0, rna.Length())
	}
	// TODO: Write test for runtime panic on dna.Postion() and dna.Range()
}

func TestRnaIupacHasMethods(t *testing.T) {
	s := new(RnaIupac)
	t.Run("Has Reverse method", func(t *testing.T) {
		if !reflect.ValueOf(s).MethodByName("Reverse").IsValid() {
			t.Error("Missing Reverse method")
		}
	})
	t.Run("Has Complement method", func(t *testing.T) {
		if !reflect.ValueOf(s).MethodByName("Complement").IsValid() {
			t.Error("Missing Complement method")
		}
	})
	t.Run("Has RevComp method", func(t *testing.T) {
		if !reflect.ValueOf(s).MethodByName("RevComp").IsValid() {
			t.Error("Missing RevComp method")
		}
	})
	t.Run("Has Alphabet method", func(t *testing.T) {
		if !reflect.ValueOf(s).MethodByName("Alphabet").IsValid() {
			t.Error("Missing Alphabet method")
		}
	})
}

func TestRnaIupacMethodsReturnTypes(t *testing.T) {
	s := new(RnaIupac)
	t.Run("Reverse returns *RnaIupac", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Reverse").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *RnaIupac")
			}
		}
	})
	t.Run("Reverse returns *RnaIupac", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Complement").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *RnaIupac")
			}
		}
	})
	t.Run("Reverse returns *RnaIupac", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("RevComp").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *RnaIupac")
			}
		}
	})
	t.Run("Alphabet returns *alphabet.RnaIupac", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Alphabet").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(new(alphabet.RnaIupac)) {
				t.Errorf("Want: %v, Got: %v",
					reflect.TypeOf(new(alphabet.RnaIupac)),
					r[i].Type(),
				)
			}
		}
	})
}

func TestRnaIupacCreation(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("RnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaIupacLetters),
				)
				dna := NewRnaIupac(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, seqLen),
		),
	)
	properties.Property("RnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaIupacLetters),
				)
				dna := NewRnaIupac(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestRnaIupacPersistence(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaIupacLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaIupacLetters),
				)
				original := NewRnaIupac(s)
				clone := new(RnaIupac)
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
					[]rune(alphabet.RnaIupacLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaIupacLetters),
				)
				original := NewRnaIupac(s)
				clone := new(RnaIupac)
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

func TestRnaIupacAccumulatesErrors(t *testing.T) {
	var _ helpers.ErrorAccumulator = new(RnaIupac)
	var seqLen uint = 1000
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}
