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
	_ sequence.Interface = new(DnaIupac)
)

func TestInitializedDnaIupac(t *testing.T) {
	dna := new(DnaIupac)

	if dna.Alphabet() != new(alphabet.DnaIupac) {
		t.Errorf("Want: %t, Got: %t", new(alphabet.DnaIupac), dna.Alphabet())
	}
	if dna.Length() != 0 {
		t.Errorf("Want: %d, Got: %d", 0, dna.Length())
	}
	// TODO: Write test for runtime panic on dna.Postion() and dna.Range()
}

func TestDnaIupacHasMethods(t *testing.T) {
	s := new(DnaIupac)
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

func TestDnaIupacMethodsReturnTypes(t *testing.T) {
	s := new(DnaIupac)
	t.Run("Reverse returns *DnaIupac", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Reverse").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *DnaIupac")
			}
		}
	})
	t.Run("Reverse returns *Rna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Complement").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *DnaIupac")
			}
		}
	})
	t.Run("Reverse returns *Rna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("RevComp").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *DnaIupac")
			}
		}
	})
	t.Run("Alphabet returns *alphabet.DnaIupac", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Alphabet").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(new(alphabet.DnaIupac)) {
				t.Errorf("Want: %v, Got: %v",
					reflect.TypeOf(new(alphabet.DnaIupac)),
					r[i].Type(),
				)
			}
		}
	})
}

func TestDnaIupacCreation(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("DnaIupac is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
				)
				dna := NewDnaIupac(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, seqLen),
		),
	)
	properties.Property("DnaIupac has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
				)
				dna := NewDnaIupac(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacPersistence(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
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
					[]rune(alphabet.DnaIupacLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
				*clone = *original
				mut := original.WithRange(n*(1/4), n*(3/4), t)
				return reflect.DeepEqual(original, clone) &&
					!reflect.DeepEqual(original, mut)
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("Reverse does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
				*clone = *original
				mut := original.Reverse()
				return reflect.DeepEqual(original, clone) &&
					!reflect.DeepEqual(original, mut)
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("Complement does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
				*clone = *original
				mut := original.Complement()
				return reflect.DeepEqual(original, clone) &&
					!reflect.DeepEqual(original, mut)
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("RevComp does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
				)
				original := NewDnaIupac(s)
				clone := new(DnaIupac)
				*clone = *original
				mut := original.RevComp()
				return reflect.DeepEqual(original, clone) &&
					!reflect.DeepEqual(original, mut)
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacAccumulatesErrors(t *testing.T) {
	var _ helpers.ErrorAccumulator = new(DnaIupac)
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}
