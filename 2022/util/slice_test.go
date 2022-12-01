package util_test

import (
	"aoc-2022/util"
	"testing"
)

func TestSortIntSliceDecending(t *testing.T) {
	tcs := []struct {
		name string
		inp  []int
		exp  []int
	}{
		{"Normal", []int{5, 3, 4, 1, 2}, []int{5, 4, 3, 2, 1}},
		{"Sorted", []int{1, 2, 3}, []int{3, 2, 1}},
		{"Empty", []int{}, []int{}},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := make([]int, len(tc.inp))
			copy(res, tc.inp)
			util.SortIntSliceDecending(res)

			if len(res) != len(tc.exp) {
				t.Fatalf("Expected length %d, got %d", len(tc.exp), len(res))
			}

			for i, v := range tc.exp {
				if res[i] != v {
					t.Fatalf("Expected %d at index %d, got %d", v, i, res[i])
				}
			}
		})
	}
}
