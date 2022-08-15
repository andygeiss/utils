package channels_test

import (
	"testing"

	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/channels"
)

func Test_Merge(t *testing.T) {
	ch1 := channels.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	ch2 := channels.Generate(13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24)
	out := channels.Merge(ch1, ch2)
	res := &mockResult{Num: 0, Sum: 0}
	for val := range out {
		res.Num++
		res.Sum += val
	}
	assert.That("number of values should be 24", t, res.Num, 24)
	assert.That("sum of the values should be 300", t, res.Sum, 300)
}
