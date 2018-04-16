package bigr

import (
	"bytes"
	"encoding/gob"
)

// Clone does a deep copy from one src to one dest
func Clone(src, dest interface{}) {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(src)
	dec.Decode(dest)
}
