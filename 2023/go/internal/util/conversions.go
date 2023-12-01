package util

import (
	"log"
	"strconv"
)

// GetIntFromString converts a string to an integer
func GetIntFromString(val string) int {
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Cannot convert string %s to int\n", val)
	}
	return valInt
}

// DigitToString converts an integer digit to a string
// Note: Only supports 0-9
func DigitToString(d int) string {
	switch d {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	}
	log.Fatalf("value %d is not a digit", d)
	return ""
}
