package channels

import (
	"runtime"
	"sync"
)

// Process works on each value from an input channel and
// sends the result to an output channel.
func Process[T any](in <-chan T, fn func(in T) T) <-chan T {
	out := make(chan T)
	// create a goroutine for each available cpu.
	num := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			for val := range in {
				// process each value
				out <- fn(val)
			}
			wg.Done()
		}()
	}
	// create a gorouting for handling the close.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
