package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay9Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day9-example.txt"), 114},
		{"Actual", util.ReadFileLines("../../inputs/day9-actual.txt"), 1934898178},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day9Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay9Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day9-example.txt"), 2},
		{"Actual", util.ReadFileLines("../../inputs/day9-actual.txt"), 1129},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day9Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
