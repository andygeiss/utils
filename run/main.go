package run

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
	// wait for close.
	for fn := range caller {
		fn()
	}
}
