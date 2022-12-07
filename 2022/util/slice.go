package util

import "sort"

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
