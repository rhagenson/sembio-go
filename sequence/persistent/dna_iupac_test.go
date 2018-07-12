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
	dna := NewDnaIupac("")

	t.Run("proper alphabet", func(t *testing.T) {
		if dna.Alphabet() != new(alphabet.DnaIupac) {
			t.Errorf("Want: %t, Got: %t", new(alphabet.DnaIupac), dna.Alphabet())
		}
	})
	t.Run("proper length", func(t *testing.T) {
		if dna.Length() != 0 {
			t.Errorf("Want: %d, Got: %d", 0, dna.Length())
		}
	})
	t.Run("proper position", func(t *testing.T) {
		if dna.Position(0) != "" {
			t.Errorf("Want: %s, Got: %s", "", dna.Position(0))
		}
	})
	t.Run("proper range", func(t *testing.T) {
		if dna.Range(0, 1) != "" {
			t.Errorf("Want: %s, Got: %s", "", dna.Range(0, 1))
		}
	})
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
	s := NewDnaIupac("")

	t.Run("Reverse returns *DnaIupac", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Reverse").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *DnaIupac")
			}
		}
	})
	t.Run("Complement returns *Rna", func(t *testing.T) {
		r := reflect.ValueOf(s).MethodByName("Complement").Call(nil)
		for i := range r {
			if r[i].Type() != reflect.TypeOf(s) {
				t.Error("Does not return a new *DnaIupac")
			}
		}
	})
	t.Run("RevComp returns *Rna", func(t *testing.T) {
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
				_ = original.Reverse()
				return reflect.DeepEqual(original, clone)
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
				_ = original.Complement()
				return reflect.DeepEqual(original, clone)
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
				_ = original.RevComp()
				return reflect.DeepEqual(original, clone)
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("start > stop errors",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}

func TestDnaIupacParallelOperations(t *testing.T) {
	var seqLen uint = 1000
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	properties.Property("NewDnaIupac(s) == NewDnaIupac(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewDnaIupac(s).Reverse() == NewDnaIupac(s).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewDnaIupac(s).RevComp() == NewDnaIupac(s).RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.Property("NewDnaIupac(s).Complement() == NewDnaIupac(s).Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := bigr.RandomStringFromRunes(
					bigr.TestSeed,
					n,
					[]rune(alphabet.DnaIupacLetters),
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
			gen.UIntRange(1, seqLen), // Length of sequence
		),
	)
	properties.TestingRun(t)
}
