package soa_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/soa"
	"os"
	"testing"
	"unsafe"
)

func Test_Allocate(t *testing.T) {
	slice := soa.Allocate[int32]()
	sizeOfOneValue := int(unsafe.Sizeof(int32(0)))
	assert.That("allocated size of [slice] is equal to page sizeOfOneValue", t, sizeOfOneValue*len(slice), os.Getpagesize())
}
