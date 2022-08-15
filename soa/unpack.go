package soa

import (
	"os"
	"unsafe"
)

// Unpack grows the slice s to the next multiple of page size.
func Unpack[T any](src []T) (dst []T) {
	tmp := make([]T, 1)
	size := os.Getpagesize() / int(unsafe.Sizeof(tmp[0]))
	next := (len(src)/size+1)*size - len(src)
	return append(src, make([]T, next, next)...)
}
