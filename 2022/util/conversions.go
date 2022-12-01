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
