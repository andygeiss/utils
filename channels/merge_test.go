package channels_test

import (
	"testing"

	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/channels"
)

func Test_Merge(t *testing.T) {
	ch1 := channels.Generate(1, 2, 3, 4)
	ch2 := channels.Generate(5, 6, 7, 8)
	out := channels.Merge(ch1, ch2)
	num := 0
	sum := 0
	for val := range out {
		sum += val
		num++
	}
	assert.That("number of values should be 24", t, num, 24)
	assert.That("sum of the values should be 300", t, sum, 300)
}
