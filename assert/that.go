package assert

import (
	"fmt"
	"testing"
)

// That compares the given values and
// fails the test if they are both not equal.
func That(desc string, value, expected interface{}, t *testing.T) {
	if fmt.Sprintf("%v", value) != fmt.Sprintf("%v", expected) {
		t.Fatalf("%s, but value is [%v]", desc, value)
	}
}
