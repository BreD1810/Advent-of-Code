package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay16Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example 1", util.ReadFileLines("../../inputs/day16-example.txt"), 46},
		{"Actual", util.ReadFileLines("../../inputs/day16-actual.txt"), 7067},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day16Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay16Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day16-example.txt"), 51},
		{"Actual", util.ReadFileLines("../../inputs/day16-actual.txt"), 7324},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day16Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
