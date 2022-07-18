package channels

import "sync"

// Merge sends each value from a list of input channels to the out channel.
func Merge[T any](in ...<-chan T) <-chan T {
	out := make(chan T)
	// create a goroutine for each input channel.
	var wg sync.WaitGroup
	wg.Add(len(in))
	for _, ch := range in {
		go func(c <-chan T) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}
	// create a gorouting for handling the close.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
