package alphabet

type StrictDNA struct{}

func (d *StrictDNA) Letters() []Letter {
	return []Letter{"A", "T", "G", "C"}
}

// Valid checks that a given Letter is in the Alphabet
func (d *StrictDNA) Valid(l Letter) (valid bool) {
	valid = false
	for _, c := range d.Letters() {
		if l == c {
			valid = true
		}
	}
	return
}

func (d *StrictDNA) Length() int {
	return len(d.Letters())
}

func (d *StrictDNA) Gapped() bool {
	return false
}

func (d *StrictDNA) Ambiguous() bool {
	return false
}

type IupacDNA struct{}

func (d *IupacDNA) Letters() []Letter {
	return []Letter{
		"A", "T", "G", "C", // Any of one nucelotide codes (i.e., 4 choose 1)
		"R", "Y", "S", "W", "K", "M", // Any of two nucelotide codes (i.e., 4 choose 2)
		"B", "D", "H", "V", "N", // Any of three nucleotide codes (i.e., 4 choose 3)
		"-", // Gap code (i.e., 4 choose 0)
	}
}

// Valid checks that a given Letter is in the Alphabet
func (d *IupacDNA) Valid(l Letter) (valid bool) {
	valid = false
	for _, c := range d.Letters() {
		if l == c {
			valid = true
		}
	}
	return
}

func (d *IupacDNA) Length() int {
	return len(d.Letters())
}

func (d *IupacDNA) Gapped() bool {
	return true
}

func (d *IupacDNA) Ambiguous() bool {
	return true
}
