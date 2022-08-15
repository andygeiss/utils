package channels_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/channels"
	"testing"
)

func Test_Multiplex(t *testing.T) {
	in := channels.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24)
	outs := channels.Multiplex(in, 2)
	res := &mockResult{Num: 0, Sum: 0}
	channels.Drain(outs[0], func(val int) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.Num++
		res.Sum += val
	})
	channels.Drain(outs[1], func(val int) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.Num++
		res.Sum += val
	})
	for {
		res.mutex.Lock()
		if res.Num == 48 {
			break
		}
		res.mutex.Unlock()
	}
	assert.That("sum of the values should be 600", t, res.Sum, 600)
}
