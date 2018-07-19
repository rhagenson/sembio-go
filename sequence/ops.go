package sequence

import "bitbucket.org/rhagenson/bigr/helpers/complement"

// SeqFunc is a type alias for a function which
// transforms a Sequence in some manner
type SeqFunc func(*Sequence) (*Sequence, error)

func reverseDna(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewDna(string(t))
}

func revCompDna(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.Atgc(t[l-1-i]), complement.Atgc(t[i])

	}
	return NewDna(string(t))
}

func complementDna(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.Atgc(byte(x.seq[i]))
	}
	return NewDna(string(t))
}

func reverseRna(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewRna(string(t))
}

func revCompRna(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.Augc(t[l-1-i]), complement.Augc(t[i])

	}
	return NewRna(string(t))
}

func complementRna(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.Augc(byte(x.seq[i]))
	}
	return NewRna(string(t))
}

func reverseDnaIupac(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewDnaIupac(string(t))
}

func revCompDnaIupac(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.DnaIupac(t[l-1-i]), complement.DnaIupac(t[i])
	}
	return NewDnaIupac(string(t))
}

func complementDnaIupac(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.DnaIupac(t[i])
	}
	return NewDnaIupac(string(t))
}

func reverseRnaIupac(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = t[l-1-i], t[i]
	}
	return NewRnaIupac(string(t))
}

func revCompRnaIupac(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l/2; i++ {
		t[i], t[l-1-i] = complement.RnaIupac(t[l-1-i]), complement.RnaIupac(t[i])

	}
	return NewRnaIupac(string(t))
}

func complementRnaIupac(x *Sequence) (*Sequence, error) {
	l := x.Length()
	t := []byte(x.seq)
	for i := uint(0); i < l; i++ {
		t[i] = complement.RnaIupac(t[i])
	}
	return NewRnaIupac(string(t))
}
