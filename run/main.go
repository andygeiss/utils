package run

import (
	"os"
	"os/signal"
	"syscall"
)

// Main ensures guarded calls to the main OS thread.
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
	// set up a signal handler.
	sigTerm := make(chan os.Signal, 2)
	signal.Notify(sigTerm, os.Interrupt, syscall.SIGTERM)
	// wait for close.
	for {
		select {
		case fn := <-caller:
			// prevent nil pointer dereference
			if fn == nil {
				return
			}
			fn()
		case <-sigTerm:
			return
		}
	}
}
