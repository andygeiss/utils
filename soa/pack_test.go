package soa_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/soa"
	"testing"
)

func Test_Pack(t *testing.T) {
	src := soa.Allocate[int32]()
	states := make([]uint64, len(src))
	states[0] = 1
	states[1] = 2
	states[111] = 3
	states[999] = 4
	packed := soa.Pack(src, states)
	assert.That("pack should shrink [packed] to 4", t, len(packed), 4)
}
