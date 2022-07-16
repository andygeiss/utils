package utils

import (
	"fmt"
	"testing"
)

// Assert helps to write better test code.
func Assert(that string, value, expected interface{}, t *testing.T) {
	if fmt.Sprintf("%v", value) != fmt.Sprintf("%v", expected) {
		t.Fatalf("%s, but value is [%v]", that, value)
	}
}
