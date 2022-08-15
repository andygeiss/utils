package channels

import "sync"

// Split sends each value from an input channel to a specific number of output channels.
func Split[T any](in <-chan T, num int) (out []chan T) {
	out = make([]chan T, num)
	for i := 0; i < num; i++ {
		out[i] = make(chan T)
	}
	// create a goroutine for each output channel.
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			for val := range in {
				out[i] <- val
			}
		}(i)
	}
	// create a goroutine for handling the close.
	go func() {
		wg.Wait()
		for i := 0; i < num; i++ {
			close(out[i])
		}
	}()
	return out
}
