package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay20Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example 1", util.ReadFileLines("../../inputs/day20-example-1.txt"), 32000000},
		{"Example 2", util.ReadFileLines("../../inputs/day20-example-2.txt"), 11687500},
		{"Actual", util.ReadFileLines("../../inputs/day20-actual.txt"), 731517480},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day20Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay20Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day20-example-2.txt"), 1},
		{"Actual", util.ReadFileLines("../../inputs/day20-actual.txt"), 244178746156661},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day20Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
