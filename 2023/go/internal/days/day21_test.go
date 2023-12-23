package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay21Part1(t *testing.T) {
	tcs := []struct {
		name      string
		inp       []string
		stepCount int
		exp       int
	}{
		{"Example", util.ReadFileLines("../../inputs/day21-example.txt"), 6, 16},
		{"Actual", util.ReadFileLines("../../inputs/day21-actual.txt"), 64, 3532},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day21Part1(tc.inp, tc.stepCount)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay21Part2(t *testing.T) {
	tcs := []struct {
		name      string
		inp       []string
		stepCount int
		exp       int
	}{
		{"Example 1", util.ReadFileLines("../../inputs/day21-example.txt"), 6, 16},
		{"Example 2", util.ReadFileLines("../../inputs/day21-example.txt"), 10, 50},
		{"Actual", util.ReadFileLines("../../inputs/day21-actual.txt"), 26501365, 590104708070703},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day21Part2(tc.inp, tc.stepCount)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
