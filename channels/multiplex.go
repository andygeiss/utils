package channels

import "sync"

// Multiplex sends each value from an input channel to every output channel.
func Multiplex[T any](in <-chan T, num int) (out []chan T) {
	out = make([]chan T, num)
	for i := 0; i < num; i++ {
		out[i] = make(chan T)
	}
	// create a goroutine for each output channel.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range in {
			for i := 0; i < num; i++ {
				out[i] <- val
			}
		}
	}()
	// create a goroutine for handling the close.
	go func() {
		wg.Wait()
		for i := 0; i < num; i++ {
			close(out[i])
		}
	}()
	return out
}
