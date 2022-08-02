package channels_test

import (
	"testing"

	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/channels"
)

func Test_Drain(t *testing.T) {
	in := channels.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24)
	num := 0
	sum := 0
	channels.Drain(in, func(val int) {
		sum += val
		num++
	})
	assert.That("number of values should be 24", num, 24, t)
	assert.That("sum of the values should be 300", sum, 300, t)
}
