package soa_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/soa"
	"os"
	"testing"
	"unsafe"
)

func Test_Unpack(t *testing.T) {
	src := soa.Allocate[int32]()
	states := make([]uint64, len(src))

	states[0] = 1
	states[1] = 2
	states[111] = 3
	states[999] = 4

	sizeOfOneValue := int(unsafe.Sizeof(int32(0)))
	packed := soa.Pack(src, states)
	unpacked := soa.Unpack(packed)

	src2 := soa.Grow(unpacked)
	src2[sizeOfOneValue+13] = 8
	packed2 := soa.Pack(src2, states)
	unpacked2 := soa.Unpack(packed2)

	assert.That("pack should shrink [packed] to 4", t, len(packed), 4)
	assert.That("unpack should grow [d] to page size", t, len(unpacked)*sizeOfOneValue, os.Getpagesize())
	assert.That("src2 should be page size 2x", t, len(src2)*sizeOfOneValue, os.Getpagesize()*2)
	assert.That("pack2 should shrink [packed] to 5", t, len(packed2), 5)
	assert.That("unpack2 should grow [src2] to page size", t, len(unpacked2)*sizeOfOneValue, os.Getpagesize())
}
