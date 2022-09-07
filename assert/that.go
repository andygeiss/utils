package assert

import (
	"fmt"
	"testing"
)

// That compares the given values and
// fails the test if they are both not equal.
func That[T any](desc string, t *testing.T, value, expected T) {
	if fmt.Sprintf("%v", value) != fmt.Sprintf("%v", expected) {
		t.Fatalf("%s, but value is [%v] instead of [%v]", desc, value, expected)
	}
}
