package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay18Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int64
	}{
		{"Example", util.ReadFileLines("../../inputs/day18-example.txt"), 62},
		{"Actual", util.ReadFileLines("../../inputs/day18-actual.txt"), 52231},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day18Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay18Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int64
	}{
		{"Example", util.ReadFileLines("../../inputs/day18-example.txt"), 952408144115},
		{"Actual", util.ReadFileLines("../../inputs/day18-actual.txt"), 57196493937398},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day18Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
