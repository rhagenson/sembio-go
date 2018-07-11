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
	_ Interface = new(DnaPersistent)
)

func TestInitializedDnaPersistent(t *testing.T) {
	dna := new(DnaPersistent)

	if dna.Alphabet() != new(alphabet.DnaStrict) {
		t.Errorf("Want: %t, Got: %t", new(alphabet.DnaStrict), dna.Alphabet())
	}
	if dna.Length() != 0 {
		t.Errorf("Want: %d, Got: %d", 0, dna.Length())
	}
	// TODO: Write test for runtime panic on dna.Postion() and dna.Range()
}

func TestDnaPersistentHasMethods(t *testing.T) {
	s := new(DnaPersistent)
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

func TestDnaPersistentMethodsReturnTypes(t *testing.T) {
	s := new(DnaPersistent)
	t.Run("Reverse returns *DnaPersistent", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Reverse").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *DnaPersistent")
			}
		}
	})
	t.Run("Reverse returns *RnaPersistent", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Complement").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *DnaPersistent")
			}
		}
	})
	t.Run("Reverse returns *RnaPersistent", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("RevComp").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *DnaPersistent")
			}
		}
	})
	t.Run("Alphabet returns *alphabet.DnaStrict", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Alphabet").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(new(alphabet.DnaStrict)) {
				t.Errorf("Want: %v, Got: %v",
					reflect.TypeOf(new(alphabet.DnaStrict)),
					r[i].Type(),
				)
			}
		}
	})
}

func TestDnaPersistentCreation(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("DnaPersistent is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaStrictLetters),
				)
				dna := NewDnaPersistent(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, seqLen),
		),
	)
	properties.Property("DnaPersistent has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaStrictLetters),
				)
				dna := NewDnaPersistent(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestDnaPersistentPersistence(t *testing.T) {
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
				original := NewDnaPersistent(s)
				clone := new(DnaPersistent)
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
				original := NewDnaPersistent(s)
				clone := new(DnaPersistent)
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
					[]rune(alphabet.DnaStrictLetters),
				)
				original := NewDnaPersistent(s)
				clone := new(DnaPersistent)
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
					[]rune(alphabet.DnaStrictLetters),
				)
				original := NewDnaPersistent(s)
				clone := new(DnaPersistent)
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
					[]rune(alphabet.DnaStrictLetters),
				)
				original := NewDnaPersistent(s)
				clone := new(DnaPersistent)
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
