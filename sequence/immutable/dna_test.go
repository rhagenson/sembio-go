package immutable_test

import (
	"strings"
	"testing"

	"bitbucket.org/rhagenson/bio"
	"bitbucket.org/rhagenson/bio/alphabet"
	"bitbucket.org/rhagenson/bio/data/codon"
	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/immutable"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestInitializedDna(t *testing.T) {
	s, _ := immutable.NewDna("")
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
	s, _ := immutable.NewDna("")

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
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Dna is same length as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := immutable.NewDna(s)
				return seq.Length() == n
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna has same positions as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := immutable.NewDna(s)
				got, _ := seq.Range(0, n)
				return got == s
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna has same internal range as input",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := immutable.NewDna(s)
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
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := immutable.NewDna(s)
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
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := immutable.NewDna(s)
				trans, _ := seq.Transcribe()

				return seq.Length() == trans.Length()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Dna is 3x length as translation",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := immutable.NewDna(s)
				trans, _ := seq.Translate(codon.Standard{}, '*')
				return seq.Length()/3 == trans.Length()
			},
			gen.UIntRange(3, sequence.TestableLength),
		),
	)

	properties.TestingRun(t)
}

func TestDnaPersistence(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("WithPosition does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				t := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				original, _ := immutable.NewDna(s)
				clone := new(immutable.Dna)
				*clone = *original
				original.With(immutable.PositionAs(n*(1/2), t))
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("WithRange does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				t := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				original, _ := immutable.NewDna(s)
				clone := new(immutable.Dna)
				*clone = *original
				original.With(immutable.RangeAs(n*(1/4), n*(3/4), t))
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Reverse does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				original, _ := immutable.NewDna(s)
				clone := new(immutable.Dna)
				*clone = *original
				original.Reverse()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				original, _ := immutable.NewDna(s)
				clone := new(immutable.Dna)
				*clone = *original
				original.Complement()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp does not mutate in-place",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				original, _ := immutable.NewDna(s)
				clone := new(immutable.Dna)
				*clone = *original
				original.RevComp()
				return original.String() == clone.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaMethodComplements(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Reverse().Reverse() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				want, _ := immutable.NewDna(s)
				rev, _ := want.Reverse()
				got, _ := rev.(*immutable.Dna).Reverse()
				return want.String() == got.(*immutable.Dna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("Complement().Complement() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				want, _ := immutable.NewDna(s)
				rev, _ := want.Complement()
				got, _ := rev.(*immutable.Dna).Complement()
				return want.String() == got.(*immutable.Dna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("RevComp().RevComp() is original",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				want, _ := immutable.NewDna(s)
				rev, _ := want.RevComp()
				got, _ := rev.(*immutable.Dna).RevComp()
				return want.String() == got.(*immutable.Dna).String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.TestingRun(t)
}

func TestDnaErrors(t *testing.T) {
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("Giving invalid input adds an error",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune("XNQZ"),
				)
				if _, err := immutable.NewDna(s); err != nil {
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
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				seq, _ := immutable.NewDna(s)
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
	parameters := gopter.DefaultTestParametersWithSeed(bio.TestSeed)
	properties := gopter.NewProperties(parameters)

	properties.Property("immutable.NewDna(s) == immutable.NewDna(s)",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				ret := make(chan *immutable.Dna)
				go func(s string, out chan *immutable.Dna) {
					seq, _ := immutable.NewDna(s)
					out <- seq
				}(s, ret)
				go func(s string, out chan *immutable.Dna) {
					seq, _ := immutable.NewDna(s)
					out <- seq
				}(s, ret)
				first := <-ret
				second := <-ret
				return first.String() == second.String()
			},
			gen.UIntRange(1, sequence.TestableLength),
		),
	)
	properties.Property("seq.(*immutable.Dna).Reverse() == seq.(*immutable.Dna).Reverse()",
		prop.ForAll(
			func(n uint) bool {
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				ret := make(chan *immutable.Dna)
				seq, _ := immutable.NewDna(s)
				go func(seq *immutable.Dna, out chan *immutable.Dna) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.Dna)
				}(seq, ret)
				go func(seq *immutable.Dna, out chan *immutable.Dna) {
					rev, _ := seq.Reverse()
					out <- rev.(*immutable.Dna)
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
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				ret := make(chan *immutable.Dna)
				seq, _ := immutable.NewDna(s)
				go func(seq *immutable.Dna, out chan *immutable.Dna) {
					rev, _ := seq.RevComp()
					out <- rev.(*immutable.Dna)
				}(seq, ret)
				go func(seq *immutable.Dna, out chan *immutable.Dna) {
					rev, _ := seq.RevComp()
					out <- rev.(*immutable.Dna)
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
				s := bio.RandomStringFromRunes(
					bio.TestSeed,
					n,
					[]rune(alphabet.Dna.String()),
				)
				ret := make(chan *immutable.Dna)
				seq, _ := immutable.NewDna(s)
				go func(seq *immutable.Dna, out chan *immutable.Dna) {
					rev, _ := seq.Complement()
					out <- rev.(*immutable.Dna)
				}(seq, ret)
				go func(seq *immutable.Dna, out chan *immutable.Dna) {
					rev, _ := seq.Complement()
					out <- rev.(*immutable.Dna)
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
