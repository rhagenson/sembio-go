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
	_ sequence.Interface = new(Dna)
)

func TestInitializedDna(t *testing.T) {
	s := NewDna("")

	t.Run("proper alphabet", func(t *testing.T) {
		if s.Alphabet() != new(alphabet.Dna) {
			t.Errorf("Want: %t, Got: %t", new(alphabet.Dna), s.Alphabet())
		}
	})
	t.Run("proper length", func(t *testing.T) {
		if s.Length() != 0 {
			t.Errorf("Want: %d, Got: %d", 0, s.Length())
		}
	})
	t.Run("proper position", func(t *testing.T) {
		if s.Position(0) != "" {
			t.Errorf("Want: %s, Got: %s", "", s.Position(0))
		}
	})
	t.Run("proper range", func(t *testing.T) {
		if s.Range(0, 1) != "" {
			t.Errorf("Want: %s, Got: %s", "", s.Range(0, 1))
		}
	})
}

func TestDnaHasMethods(t *testing.T) {
	s := NewDna("")

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

func TestDnaMethodsReturnTypes(t *testing.T) {
	s := NewDna("")

	t.Run("Reverse returns *Dna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Reverse").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *Dna")
			}
		}
	})
	t.Run("Complement returns *Dna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Complement").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *Dna")
			}
		}
	})
	t.Run("RevComp returns *Dna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("RevComp").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *Dna")
			}
		}
	})
	t.Run("Alphabet returns *alphabet.DnaStrict", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Alphabet").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(new(alphabet.Dna)) {
				t.Errorf("Want: %v, Got: %v",
					reflect.TypeOf(new(alphabet.Dna)),
					r[i].Type(),
				)
			}
		}
	})
}

func TestDnaCreation(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Dna is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
				)
				dna := NewDna(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, seqLen),
		),
	)
	properties.Property("Dna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
				)
				dna := NewDna(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestDnaPersistence(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
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
					[]rune(alphabet.DnaLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
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
					[]rune(alphabet.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
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
					[]rune(alphabet.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
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
					[]rune(alphabet.DnaLetters),
				)
				original := NewDna(s)
				clone := new(Dna)
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

func TestDnaAccumulatesErrors(t *testing.T) {
	var _ helpers.ErrorAccumulator = new(Dna)
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
				seq := NewDna(s)
				for _, err := range seq.errs {
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("start > stop errors",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
				)
				seq := NewDna(s)
				seq.Range(n, 0)
				for _, err := range seq.errs {
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestDnaParallelOperations(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewDna(s) == NewDna(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewDna(s).Reverse() == NewDna(s).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewDna(s).RevComp() == NewDna(s).RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewDna(s).Complement() == NewDna(s).Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}
