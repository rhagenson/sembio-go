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
	_ sequence.Interface = new(Rna)
)

func TestInitializedRna(t *testing.T) {
	s := NewRna("")

	t.Run("proper alphabet", func(t *testing.T) {
		if s.Alphabet() != new(alphabet.Rna) {
			t.Errorf("Want: %t, Got: %t", new(alphabet.Rna), s.Alphabet())
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

func TestRnaHasMethods(t *testing.T) {
	s := NewRna("")

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

func TestRnaMethodsReturnTypes(t *testing.T) {
	s := NewRna("")

	t.Run("Reverse returns *Rna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Reverse").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *Rna")
			}
		}
	})
	t.Run("Complement returns *Rna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Complement").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *Rna")
			}
		}
	})
	t.Run("RevComp returns *Rna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("RevComp").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *Rna")
			}
		}
	})
	t.Run("Alphabet returns *alphabet.RnaStrict", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Alphabet").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(new(alphabet.Rna)) {
				t.Errorf("Want: %v, Got: %v",
					reflect.TypeOf(new(alphabet.Rna)),
					r[i].Type(),
				)
			}
		}
	})
}

func TestRnaCreation(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("Rna is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
				)
				dna := NewRna(s)
				return dna.Length() == n
			},
			gen.UIntRange(1, seqLen),
		),
	)
	properties.Property("Rna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
				)
				dna := NewRna(s)
				got := dna.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestRnaPersistence(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
				)
				original := NewRna(s)
				clone := new(Rna)
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
					[]rune(alphabet.RnaLetters),
				)
				t := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
				)
				original := NewRna(s)
				clone := new(Rna)
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

func TestRnaAccumulatesErrors(t *testing.T) {
	var _ helpers.ErrorAccumulator = new(Rna)
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
				seq := NewRna(s)
				for _, err := range seq.errs {
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("start > stop errors",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
				)
				seq := NewRna(s)
				seq.Range(n, 0)
				for _, err := range seq.errs {
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestRnaParallelOperations(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewRna(s) == NewRna(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewRna(s).Reverse() == NewRna(s).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewRna(s).RevComp() == NewRna(s).RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewRna(s).Complement() == NewRna(s).Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.RnaLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}
