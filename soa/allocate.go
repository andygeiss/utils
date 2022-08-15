package soa

import (
	"os"
	"unsafe"
)

// Allocate creates a new slice s by using a total size of a page size.
// The length and capacity is set to "page size / size of type".
func Allocate[T any]() (t []T) {
	tmp := make([]T, 1)
	size := os.Getpagesize() / int(unsafe.Sizeof(tmp[0]))
	return make([]T, size, size)
}
