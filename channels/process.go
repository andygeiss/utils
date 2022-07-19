package channels

import (
	"context"
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
			defer wg.Done()
			for val := range in {
				// process each value
				out <- fn(val)
			}
		}()
	}
	// create a gorouting for handling the close.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Process works on each value from an input channel and
// sends the result to an output channel.
func ProcessWithContext[T any](ctx context.Context, in <-chan T, fn func(ctx context.Context, in T) T) <-chan T {
	out := make(chan T)
	// create a goroutine for each available cpu.
	num := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				// check for cancel or timeout etc.
				case <-ctx.Done():
					return
				// delegate work to the function
				case val := <-in:
					out <- fn(ctx, val)
				}
			}
		}()
	}
	// create a goroutine for handling the close.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
