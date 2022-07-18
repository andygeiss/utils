package channels_test

import (
	"context"
	"testing"
	"time"

	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/channels"
)

func Test_Process(t *testing.T) {
	ch := channels.Generate(1, 2, 3, 4)
	fn := func(in int) (out int) {
		return in + 1
	}
	out := channels.Process(ch, fn)
	num := 0
	sum := 0
	for val := range out {
		sum += val
		num++
	}
	assert.That("number of values should be 4", num, 4, t)
	assert.That("sum of the values should be 14", sum, 14, t)
}

func Test_ProcessWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ch := channels.Generate(1, 2, 3, 4)
	fn := func(ctx context.Context, in int) (out int) {
		return in + 1
	}
	out := channels.ProcessWithContext(ctx, ch, fn)
	cancel()
	num := 0
	sum := 0
	for val := range out {
		sum += val
		num++
	}
	assert.That("number of values should be 4", num, 4, t)
	assert.That("sum of the values should be 14", sum, 14, t)
}
