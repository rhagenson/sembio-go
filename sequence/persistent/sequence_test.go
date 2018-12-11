package persistent_test

import (
	"bitbucket.org/rhagenson/bio/sequence"
	"bitbucket.org/rhagenson/bio/sequence/persistent"
)

var _ sequence.Interface = new(persistent.Struct)
var _ persistent.Wither = new(persistent.Struct)
var _ persistent.Validator = new(persistent.Struct)
