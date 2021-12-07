package util

import (
	"math"
	"strconv"
)

func PowWithInts(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func BinaryStringToDecimal(s string) int64 {
	res, _ := strconv.ParseInt(s, 2, 64)
	return res
}

func IntAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
