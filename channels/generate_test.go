package channels_test

import (
	"testing"

	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/channels"
)

func Test_Generate(t *testing.T) {
	out := channels.Generate(1, 2, 3, 4)
	num := 0
	sum := 0
	for val := range out {
		sum += val
		num++
	}
	assert.That("number of values should be 4", num, 4, t)
	assert.That("sum of the values should be 10", sum, 10, t)
}
