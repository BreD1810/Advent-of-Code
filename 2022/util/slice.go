package util

import "sort"

//SortIntSliceDecending sorts an int slice in decending order
func SortIntSliceDecending(s []int) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}
