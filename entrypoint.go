package utils

import "runtime"

// Ensures that the main goroutine is bound to the main thread.
func init() {
	runtime.LockOSThread()
}

var wrapper = func(fn func()) {
	panic("utils.Main(..) must be called before utils.Run(..)")
}

// Main ensures guarded calls to the main thread.
func Main(main func()) {
	caller := make(chan func())
	// specify a wrapper function
	wrapper = func(fn func()) {
		done := make(chan bool)
		caller <- func() {
			fn()
			done <- true
		}
		<-done
	}
	// start the main function at the main OS thread.
	go func() {
		main()
		close(caller)
	}()
	// wait for close.
	for fn := range caller {
		fn()
	}
}

// Run handles calls to "main-thread-sensitive" libraries by
// using the main OS thread only.
func Run(fn func()) {
	wrapper(fn)
}
