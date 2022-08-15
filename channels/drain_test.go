package channels_test

import (
	"sync"
	"testing"

	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/channels"
)

type mockResult struct {
	Num   int
	Sum   int
	mutex sync.Mutex
}

func Test_Drain(t *testing.T) {
	in := channels.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24)
	res := &mockResult{Num: 0, Sum: 0}
	channels.Drain(in, func(val int) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.Num++
		res.Sum += val
	})
	assert.That("number of values should be 24", t, res.Num, 24)
	assert.That("sum of the values should be 300", t, res.Sum, 300)
}
