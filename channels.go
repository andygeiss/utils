package utils

import (
	"runtime"
	"sync"
)

// Generate sends each value to the out channel.
func Generate[T any](values ...T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, val := range values {
			out <- val
		}
	}()
	return out
}

// Merge sends each value from a list of input channels to the out channel.
func Merge[T any](in ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(in))
	for _, ch := range in {
		go func(c <-chan T) {
			for val := range c {
				out <- val
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Process works on each value from an input channel and
// sends the result to an output channel.
func Process[T any](in <-chan T, fn func(in T) T) <-chan T {
	out := make(chan T)
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
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
