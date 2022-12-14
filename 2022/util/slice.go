package util

import (
	"fmt"
	"sort"
)

//SortIntSliceDecending sorts an int slice in decending order
func SortIntSliceDecending(s []int) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}

// SumIntSlice sums the integers in an []int
func SumIntSlice(s []int) int {
	var total int
	for _, i := range s {
		total += i
	}
	return total
}

// Print2DRuneArray prints a 2D rune array
func Print2DRuneArray(rs [][]rune) {
	for _, row := range rs {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}
