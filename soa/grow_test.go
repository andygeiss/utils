package soa_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/soa"
	"os"
	"testing"
	"unsafe"
)

func Test_Grow(t *testing.T) {
	slice := soa.Allocate[int32]()
	slice = soa.Grow(slice)
	sizeOfOneValue := int(unsafe.Sizeof(int32(0)))
	assert.That("grow size  of [b] is equal to 2x page size", t, sizeOfOneValue*len(slice), os.Getpagesize()*2)
}
