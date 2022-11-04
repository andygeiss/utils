package soa_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/soa"
	"testing"
)

func Test_Unpack(t *testing.T) {
	// Arrange
	src := make([]int32, 4096)
	states := make([]uint64, len(src))
	// Act
	states[0] = 1
	states[1] = 2
	states[111] = 3
	states[999] = 4
	packed := soa.Pack(src, states)
	unpacked := soa.Unpack(packed, 4096)
	// Assert
	assert.That("pack should shrink [packed] to 4", t, len(packed), 4)
	assert.That("unpack should grow [d] to 4096", t, len(unpacked), 4096)
}
