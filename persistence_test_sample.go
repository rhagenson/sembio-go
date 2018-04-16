package bigr

import (
	"reflect"
	"testing"
)

// TestPersistenceExample is a sample test for how to structure the check on a persistent type
// Integers are used a simple persistent object such that a value is assigned to a, an operation is applied on a,
// then we check that a did not chnage from its original state.
// This property should hold for all operations on any persistent type (i.e. a.Op() -> !a && !a.!Op() -> a)
func TestPersistenceExample(t *testing.T) {
	// Step 1. Assign initial value
	var initial = 1

	// Step 2. Create a deep copy of the value (for comparison later)
	var clone int
	Clone(initial, &clone)

	// Step 3. Apply an operation on initial value
	_ = initial + 1

	// Step 4. Check that the initial value and post-operation value are the same
	if !reflect.DeepEqual(initial, clone) {
		t.Errorf("%v not DeepEqual to %v", initial, clone)
	}
}
