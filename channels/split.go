package channels

import "sync"

// Split ...
func Split[T any](in <-chan T, num int) (out []chan T) {
	out = make([]chan T, num)
	for i := 0; i < num; i++ {
		out[i] = make(chan T)
	}
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
	go func() {
		wg.Wait()
		for i := 0; i < num; i++ {
			close(out[i])
		}
	}()
	return out
}
