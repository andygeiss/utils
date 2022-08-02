package run

import "runtime"

// Ensures that the main goroutine is bound to the main thread.
func init() {
	runtime.LockOSThread()
}

var wrapper = func(fn func()) {
	panic("run.Main(..) must be called before run.Safe(..)")
}
