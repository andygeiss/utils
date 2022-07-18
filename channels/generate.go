package channels

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
