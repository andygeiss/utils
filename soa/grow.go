package soa

import (
	"os"
	"unsafe"
)

// Grow appends a new page of values to a given slice src
// and save it as a new slice dst.
func Grow[T any](src []T) (out []T) {
	tmp := make([]T, 1)
	size := os.Getpagesize() / int(unsafe.Sizeof(tmp[0]))
	return append(src, make([]T, size, size)...)
}
