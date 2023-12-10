package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay10Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example 1", util.ReadFileLines("../../inputs/day10-example-1.txt"), 4},
		{"Example 2", util.ReadFileLines("../../inputs/day10-example-2.txt"), 8},
		{"Actual", util.ReadFileLines("../../inputs/day10-actual.txt"), 6909},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day10Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay10Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		//{"Example 3", util.ReadFileLines("../../inputs/day10-example-3.txt"), 4},
		//{"Example 4", util.ReadFileLines("../../inputs/day10-example-4.txt"), 4},
		//{"Example 5", util.ReadFileLines("../../inputs/day10-example-5.txt"), 8},
		{"Example 6", util.ReadFileLines("../../inputs/day10-example-6.txt"), 10},
		{"Actual", util.ReadFileLines("../../inputs/day10-actual.txt"), 461},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day10Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
