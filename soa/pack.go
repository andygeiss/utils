package soa

import "reflect"

// Pack shrinks the slice s to a total amount of values which are not zero (0).
// The order of the elements remains the same.
func Pack[T any](src []T) (dst []T) {
	// create an empty slice for the result.
	out := make([]T, 0)
	// now check each element for a value of zero(0).
	for i := 0; i < len(src); i++ {
		// if value is zero (0) then skip.
		if reflect.ValueOf(src[i]).Int() == reflect.ValueOf(0).Int() {
			continue
		}
		// add the non-zero value to the result.
		out = append(out, src[i])
	}
	// finally, return the packed slice.
	return out
}
