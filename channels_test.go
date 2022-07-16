package utils_test

import (
	"testing"

	"github.com/andygeiss/utils"
)

func Test_Generate(t *testing.T) {
	out := utils.Generate(1, 2, 3, 4)
	num := 0
	sum := 0
	for val := range out {
		sum += val
		num++
	}
	utils.Assert("number of values should be 4", num, 4, t)
	utils.Assert("sum of the values should be 10", sum, 10, t)
}

func Test_Merge(t *testing.T) {
	ch1 := utils.Generate(1, 2, 3, 4)
	ch2 := utils.Generate(5, 6, 7, 8)
	out := utils.Merge(ch1, ch2)
	num := 0
	sum := 0
	for val := range out {
		sum += val
		num++
	}
	utils.Assert("number of values should be 8", num, 8, t)
	utils.Assert("sum of the values should be 14", sum, 36, t)
}

func Test_Process(t *testing.T) {
	ch := utils.Generate(1, 2, 3, 4)
	fn := func(in int) (out int) {
		return in + 1
	}
	out := utils.Process(ch, fn)
	num := 0
	sum := 0
	for val := range out {
		sum += val
		num++
	}
	utils.Assert("number of values should be 4", num, 4, t)
	utils.Assert("sum of the values should be 14", sum, 14, t)
}
