package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay12Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day12-example.txt"), 21},
		{"Actual", util.ReadFileLines("../../inputs/day12-actual.txt"), 6935},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day12Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay12Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day12-example.txt"), 525152},
		{"Actual", util.ReadFileLines("../../inputs/day12-actual.txt"), 3920437278260},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day12Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
