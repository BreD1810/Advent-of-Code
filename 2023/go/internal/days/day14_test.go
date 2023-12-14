package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay14Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example 1", util.ReadFileLines("../../inputs/day14-example.txt"), 136},
		{"Actual", util.ReadFileLines("../../inputs/day14-actual.txt"), 113525},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day14Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay14Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day14-example.txt"), 64},
		{"Actual", util.ReadFileLines("../../inputs/day14-actual.txt"), 101292},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day14Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
