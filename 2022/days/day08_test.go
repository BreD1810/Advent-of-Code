package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay8Part1(t *testing.T) {
	exampleInp := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	tcs := []struct {
		name string
		inp  [][]int
		exp  int
	}{
		{"Example", exampleInp, 21},
		{"Actual", util.ReadFileIntLines("../inputs/day08-1.txt"), 1825},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day8Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay8Part2(t *testing.T) {
	exampleInp := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	tcs := []struct {
		name string
		inp  [][]int
		exp  int
	}{
		{"Example", exampleInp, 8},
		{"Actual", util.ReadFileIntLines("../inputs/day08-1.txt"), 235200},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day8Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
