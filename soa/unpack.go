package soa

// Unpack grows the slice s to the next multiple of size.
func Unpack[T any](src []T, size int) (dst []T) {
	next := (len(src)/size+1)*size - len(src)
	return append(src, make([]T, next, next)...)
}
