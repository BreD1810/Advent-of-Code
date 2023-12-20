package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay8Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day8-example-1.txt"), 2},
		{"Example", util.ReadFileLines("../../inputs/day8-example-2.txt"), 6},
		{"Actual", util.ReadFileLines("../../inputs/day8-actual.txt"), 19241},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day8Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay8Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day8-example-3.txt"), 6},
		{"Actual", util.ReadFileLines("../../inputs/day8-actual.txt"), 9606140307013},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day8Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
