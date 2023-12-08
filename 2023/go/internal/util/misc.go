package util

import (
	"math"
	"strconv"
)

var MaxInt = int(^uint(0) >> 1)

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

type Coordinate struct {
	X, Y int
}

type Coordinate3D struct {
	X, Y, Z int
}

func GetLeastCommonMultiple(ns []int) int {
	lcm := ns[0]
	for i := 0; i < len(ns); i++ {
		num1 := lcm
		num2 := ns[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * ns[i]) / gcd
	}

	return lcm
}
