package sequence

// WithFunc is a transformative function that can be chained
type WithFunc func(*sequence) *sequence

// Wither provides a variadic method to transform the sequence
type Wither interface {
	With(...WithFunc) *sequence
}

// PositionAs mutates a sequence position
func PositionAs(n uint, pos string) WithFunc {
	return func(x *sequence) *sequence {
		x.seq = x.seq[:n] + pos + x.seq[n+1:]
		return x
	}
}

// RangeAs mutates a sequence range
func RangeAs(st, sp uint, pos string) WithFunc {
	return func(x *sequence) *sequence {
		x.seq = x.seq[:st] + pos + x.seq[sp:]
		return x
	}
}
