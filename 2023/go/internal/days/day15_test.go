package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay15Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  string
		exp  int
	}{
		{"Example 1", "HASH", 52},
		{"Example 2", util.ReadFileLine("../../inputs/day15-example.txt"), 1320},
		{"Actual", util.ReadFileLine("../../inputs/day15-actual.txt"), 507666},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day15Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay15Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  string
		exp  int
	}{
		{"Example", util.ReadFileLine("../../inputs/day15-example.txt"), 145},
		{"Actual", util.ReadFileLine("../../inputs/day15-actual.txt"), 233537},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day15Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
