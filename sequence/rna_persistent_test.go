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

var (
	_ Interface = new(RnaPersistent)
)

func TestInitializedRnaPersistent(t *testing.T) {
	rna := new(RnaPersistent)

	if rna.Alphabet() != new(alphabet.RnaStrict) {
		t.Errorf("Want: %t, Got: %t", new(alphabet.RnaStrict), rna.Alphabet())
	}
	if rna.Length() != 0 {
		t.Errorf("Want: %d, Got: %d", 0, rna.Length())
	}
	// TODO: Write test for runtime panic on dna.Postion() and dna.Range()
}

func TestRnaPersistentHasMethods(t *testing.T) {
	s := new(RnaPersistent)
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

func TestRnaPersistentMethodsReturnTypes(t *testing.T) {
	s := new(RnaPersistent)
	t.Run("Reverse returns *RnaPersistent", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Reverse").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *RnaPersistent")
			}
		}
	})
	t.Run("Reverse returns *RnaPersistent", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Complement").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *RnaPersistent")
			}
		}
	})
	t.Run("Reverse returns *RnaPersistent", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("RevComp").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *RnaPersistent")
			}
		}
	})
	t.Run("Alphabet returns *alphabet.RnaStrict", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Alphabet").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(new(alphabet.RnaStrict)) {
				t.Errorf("Want: %v, Got: %v",
					reflect.TypeOf(new(alphabet.RnaStrict)),
					r[i].Type(),
				)
			}
		}
	})
}

func TestRnaPersistentCreation(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("RnaPersistent is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaStrictLetters),
				)
				dna := NewRnaPersistent(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, seqLen),
		),
	)
	properties.Property("RnaPersistent has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaStrictLetters),
				)
				dna := NewRnaPersistent(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestRnaPersistentPersistence(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaStrictLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaStrictLetters),
				)
				original := NewRnaPersistent(s)
				clone := new(RnaPersistent)
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
					[]rune(alphabet.RnaStrictLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaStrictLetters),
				)
				original := NewRnaPersistent(s)
				clone := new(RnaPersistent)
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
