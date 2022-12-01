package util

import (
	"math"
	"strconv"
)

// PowWithInts lets you use math.pow with integers
func PowWithInts(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// BinaryStringToDecimal converts a binary string to a integer
func BinaryStringToDecimal(s string) int64 {
	res, _ := strconv.ParseInt(s, 2, 64)
	return res
}

// IntAbs returns the absolute value of an integer
func IntAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
