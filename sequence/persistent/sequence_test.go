package persistent

import "bitbucket.org/rhagenson/bio/sequence"

var _ sequence.Interface = new(Struct)
var _ Wither = new(Struct)
var _ Validator = new(Struct)
