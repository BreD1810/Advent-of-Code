package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay18Part1(t *testing.T) {
	exampleInp := []string{
		"2,2,2",
		"1,2,2",
		"3,2,2",
		"2,1,2",
		"2,3,2",
		"2,2,1",
		"2,2,3",
		"2,2,4",
		"2,2,6",
		"1,2,5",
		"3,2,5",
		"2,1,5",
		"2,3,5",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 64},
		{"Actual", util.ReadFileLines("../inputs/day18-1.txt"), 4288},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day18Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay18Part2(t *testing.T) {
	exampleInp := []string{
		"2,2,2",
		"1,2,2",
		"3,2,2",
		"2,1,2",
		"2,3,2",
		"2,2,1",
		"2,2,3",
		"2,2,4",
		"2,2,6",
		"1,2,5",
		"3,2,5",
		"2,1,5",
		"2,3,5",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 58},
		{"Actual", util.ReadFileLines("../inputs/day18-1.txt"), 2494},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day18Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
