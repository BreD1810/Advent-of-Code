package util

import (
	"fmt"
	"strconv"
)

func GetIntFromString(val string) int {
	val_int, err := strconv.Atoi(val)
	if err != nil {
		fmt.Printf("Cannot convert %s to int", val)
	}
	return val_int
}
