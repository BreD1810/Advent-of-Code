package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay17Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example 1", util.ReadFileLines("../../inputs/day17-example.txt"), 102},
		{"Actual", util.ReadFileLines("../../inputs/day17-actual.txt"), 1195},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day17Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay17Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day17-example.txt"), 94},
		{"Actual", util.ReadFileLines("../../inputs/day17-actual.txt"), 1347},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day17Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
