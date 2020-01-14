package mutable_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sembio/go/bio/alphabet"
	"github.com/sembio/go/bio/alphabet/hashmap"
	"github.com/sembio/go/bio/data/codon"
	"github.com/sembio/go/bio/sequence"
	"github.com/sembio/go/bio/sequence/mutable"
	"github.com/sembio/go/bio/test"
	"github.com/sembio/go/bio/utils"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedDna(t *testing.T) {
	s, _ := mutable.NewDna("")
	t.Run("Length is zero", sequence.TestLengthIs(s, 0))
	t.Run("Position is empty", sequence.TestPositionIs(s, 0, ""))
	t.Run("Range is empty", sequence.TestRangeIs(s, 0, 1, ""))
	t.Run("Transcribe is empty", func(t *testing.T) {
		r, _ := s.Transcribe()
		if r.Length() != 0 {
			t.Errorf("Nucleotides gained in Transcribe()")
		}
	})
	t.Run("Translate is empty", func(t *testing.T) {
		r, _ := s.Translate(codon.Standard{}, '*')
		if r.Length() != 0 {
			t.Errorf("Amino acids gained in Translate()")
		}
	})
}

func TestDnaHasMethods(t *testing.T) {
	s, _ := mutable.NewDna("")

	t.Run("Has Reverse method", func(t *testing.T) {
		if _, err := s.Reverse(); err != nil {
			t.Error("Reverse method does not exist")
		}
	})
	t.Run("Has Complement method", func(t *testing.T) {
		if _, err := s.Complement(); err != nil {
			t.Error("Complement method does not exist")
		}
	})
	t.Run("Has RevComp method", func(t *testing.T) {
		if _, err := s.RevComp(); err != nil {
			t.Error("RevComp method does not exist")
		}
	})
	t.Run("Has Transcribe method", func(t *testing.T) {
		if _, err := s.Transcribe(); err != nil {
			t.Error("Translate method does not exist")
		}
	})
	t.Run("Has Translate method", func(t *testing.T) {
		if _, err := s.Translate(codon.Standard{}, '*'); err != nil {
			t.Error("Translate method does not exist")
		}
	})
}

func TestDnaCreation(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Dna is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				seq, _ := mutable.NewDna(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				seq, _ := mutable.NewDna(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				seq, _ := mutable.NewDna(s)
				onefourth := n / 4
				threefourths := n * 3 / 4
				got, _ := seq.Range(onefourth, threefourths)
				return got == s[onefourth:threefourths]
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna has same internal postions as input",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				seq, _ := mutable.NewDna(s)
				onefourth := n / 4
				threefourth := n * (3 / 4)
				gotoneforth, _ := seq.Position(onefourth)
				wantoneforth := string(s[onefourth])
				gotthreeforth, _ := seq.Position(threefourth)
				wantthreeforth := string(s[threefourth])
				return gotoneforth == wantoneforth && gotthreeforth == wantthreeforth
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna is same length as transcription",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				seq, _ := mutable.NewDna(s)
				trans, _ := seq.Transcribe()

				return seq.Length() == trans.Length()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna is 3x length as translation",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				seq, _ := mutable.NewDna(s)
				trans, _ := seq.Translate(codon.Standard{}, '*')
				return seq.Length()/3 == trans.Length()
			},
			gen.UIntRange(3, sequence.TestableLength),
		),
	)

	properties.TestingRun(t)
}

func TestDnaMutability(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				original, _ := mutable.NewDna(s)
				clone, _ := mutable.NewDna(s)
				original.With(mutable.PositionAs(n/2, t))
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("WithRange mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				t := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				original, _ := mutable.NewDna(s)
				clone, _ := mutable.NewDna(s)
				original.With(mutable.RangeAs(n/4, n*3/4, t))
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Reverse mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				if s == utils.ReverseByBytes(s) { // Skip palindromes
					return true
				}
				original, _ := mutable.NewDna(s)
				clone, _ := mutable.NewDna(s)
				original.Reverse()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				original, _ := mutable.NewDna(s)
				clone, _ := mutable.NewDna(s)
				original.Complement()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp mutates in-place",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				original, _ := mutable.NewDna(s)
				clone, _ := mutable.NewDna(s)
				original.RevComp()
				return original.String() != clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				want, _ := mutable.NewDna(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*mutable.Dna).Reverse()
				return want.String() == got.(*mutable.Dna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement().Complement() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				want, _ := mutable.NewDna(s)
				rev, _ := want.Complement()
				got, _ := rev.(*mutable.Dna).Complement()
				return want.String() == got.(*mutable.Dna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp().RevComp() is original",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				want, _ := mutable.NewDna(s)
				rev, _ := want.RevComp()
				got, _ := rev.(*mutable.Dna).RevComp()
				return want.String() == got.(*mutable.Dna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(string(alphabet.TestExcludesSingleLetters([]byte(hashmap.NewDna().String())))),
				)
				if _, err := mutable.NewDna(s); err != nil {
					if !strings.Contains(err.Error(), "not in alphabet") {
						t.Errorf("Dna creation error should mention not in alphabet")
						return false
					}
				} else {
					t.Errorf("Dna should error when using invalid characters, error")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("start > stop errors",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				seq, _ := mutable.NewDna(s)
				_, err := seq.Range(n, 0)
				if err == nil {
					t.Errorf("Dna should accumulate an err during Range() when start > stop")
					return false
				}
				if !strings.Contains(err.Error(), "impossible range") {
					t.Errorf("Dna Range error should mention impossible range")
					return false
				}
				return true
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaParallelOperations(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(test.Seed)
	properties := gopter.NewProperties(parameters)

	properties.Property("mutable.NewDna(s) == mutable.NewDna(s)",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				ret := make(chan *mutable.Dna)
				go func(s string, out chan *mutable.Dna) {
					seq, _ := mutable.NewDna(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *mutable.Dna) {
					seq, _ := mutable.NewDna(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.(*mutable.Dna).Reverse() == seq.(*mutable.Dna).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				ret := make(chan *mutable.Dna)
				seq, _ := mutable.NewDna(s)
				go func(seq *mutable.Dna, out chan *mutable.Dna) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.Dna)
				}(seq, ret)
				go func(seq *mutable.Dna, out chan *mutable.Dna) {
					rev, _ := seq.Reverse()
					out <- rev.(*mutable.Dna)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.RevComp() == seq.RevComp()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				ret := make(chan *mutable.Dna)
				seq, _ := mutable.NewDna(s)
				go func(seq *mutable.Dna, out chan *mutable.Dna) {
					rev, _ := seq.RevComp()
					out <- rev.(*mutable.Dna)
				}(seq, ret)
				go func(seq *mutable.Dna, out chan *mutable.Dna) {
					rev, _ := seq.RevComp()
					out <- rev.(*mutable.Dna)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.Complement() == seq.Complement()",
		prop.ForAll(
			func(n uint) bool {
				s := test.RandomStringFromRunes(
					test.Seed,
					n,
					[]rune(hashmap.NewDna().String()),
				)
				ret := make(chan *mutable.Dna)
				seq, _ := mutable.NewDna(s)
				go func(seq *mutable.Dna, out chan *mutable.Dna) {
					rev, _ := seq.Complement()
					out <- rev.(*mutable.Dna)
				}(seq, ret)
				go func(seq *mutable.Dna, out chan *mutable.Dna) {
					rev, _ := seq.Complement()
					out <- rev.(*mutable.Dna)
				}(seq, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

// Building a new Dna from valid letters results in no error
func ExampleNewDna_errorless() {
	s, err := mutable.NewDna("ATGC")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// ATGC, <nil>
}

// Building a new Dna from invalid letters results in an error
// Note that only the first error is returned, not all errors
// The invalid '%' is caught, but nothing is said of the invalid '&'
func ExampleNewDna_errored() {
	s, err := mutable.NewDna("%" + "ATGC" + "&")

	fmt.Printf("%s, %v", s, err)
	// Output:
	// %ATGC&, "%" not in alphabet
}

// Reversing a valid Dna results in no error
func ExampleDna_Reverse_errorless() {
	s, _ := mutable.NewDna("ATGC")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// CGTA, <nil>
}

// Reversing an invalid Dna results in an error
// Note that only the first error is returned, not all errors
// The invalid '&' is caught, but nothing is said of the invalid '%'
func ExampleDna_Reverse_errored() {
	s, _ := mutable.NewDna("%" + "ATGC" + "&")
	rev, err := s.Reverse()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// &CGTA%, "&" not in alphabet
}

// Reverse complementing a valid Dna results in no error
func ExampleDna_RevComp_errorless() {
	s, _ := mutable.NewDna("ATGC")
	rev, err := s.RevComp()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// GCAT, <nil>
}

// Reverse complementing an invalid Dna results in an error
// Note that both invalid letters '%' and '&' became 'X' (which is also an invalid letter)
func ExampleDna_RevComp_errored() {
	s, err := mutable.NewDna("%" + "ATGC" + "&")
	rev, err := s.RevComp()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// XGCATX, "X" not in alphabet
}

// Complementing a valid Dna results in no error
func ExampleDna_Complement_errorless() {
	s, _ := mutable.NewDna("ATGC")
	rev, err := s.Complement()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// TACG, <nil>
}

// Complementing an invalid Dna results in an error
// Note that both invalid letters '%' and '&' became 'X' (which is also an invalid letter)
func ExampleDna_Complement_errored() {
	s, err := mutable.NewDna("%" + "ATGC" + "&")
	rev, err := s.Complement()

	fmt.Printf("%s, %v", rev, err)
	// Output:
	// XTACGX, "X" not in alphabet
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewDna()
func ExampleDna_Alphabet() {
	s, _ := mutable.NewDna("ATGC")

	fmt.Println(s.Alphabet())
	// Output:
	// ACGT
}

// Note that the alphabet gets sorted and would be
// unaffected by an invalid input to mutable.NewDna()
func ExampleDna_LetterCount() {
	s, _ := mutable.NewDna("ATGC" + "AAAA")

	fmt.Println(s.LetterCount())
	// Output:
	// map[A:5 C:1 G:1 T:1]
}
