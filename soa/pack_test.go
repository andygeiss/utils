package soa_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/soa"
	"testing"
)

func Test_Pack(t *testing.T) {
	src := soa.Allocate[int32]()
	src[0] = 1
	src[1] = 2
	src[111] = 3
	src[999] = 4
	packed := soa.Pack(src)
	assert.That("pack should shrink [packed] to 4", t, len(packed), 4)
}
