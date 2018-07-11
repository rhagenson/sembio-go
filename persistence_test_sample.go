package bigr

import (
	"reflect"
	"testing"
)

// TestPersistenceExample is a sample test for how to structure the check on a persistent type
// Integers are used a simple persistent object such that a value is assigned to original, an operation is applied on original, then original is checked,
// then we check that original did not chnage state.
// This property should hold for all operations on any persistent type:
// 		original.Op() -> !original && !original.!Op() -> original
func TestPersistenceExample(t *testing.T) {
	// Step 1. Assign initial value
	var original = 1

	// Step 2. Create a deep copy of the value (for comparison later)
	// 		Note that `clone := original` simply stores two pointers so
	//		comparison would always yield an object compared to itself
	//		Second note: This DeepClone only works for structs with all
	//		exported fields, alternatively use `*clone = *original`
	var clone int
	DeepClone(&original, &clone)

	// Step 3. Apply an operation
	// 		Note here that the return value is ignored.
	//		We are checking that original does not change.
	_ = original + 1

	// Step 4. Check that the original value and post-operation value are the same
	if !reflect.DeepEqual(original, clone) {
		t.Errorf("%v not DeepEqual to %v", original, clone)
	}
}
