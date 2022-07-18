package run

// Safe handles calls to "main-thread-sensitive" libraries by
// using the main OS thread only.
func Safe(fn func()) {
	wrapper(fn)
}
