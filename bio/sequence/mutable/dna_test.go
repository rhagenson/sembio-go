package mutable_test

import (
	"strings"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/bio-ext/bio-go/bio/alphabet"
	"github.com/bio-ext/bio-go/bio/alphabet/hashmap"
	"github.com/bio-ext/bio-go/bio/data/codon"
	"github.com/bio-ext/bio-go/bio/sequence"
	"github.com/bio-ext/bio-go/bio/sequence/mutable"
	"github.com/bio-ext/bio-go/bio/test"
	"github.com/bio-ext/bio-go/bio/utils"
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
