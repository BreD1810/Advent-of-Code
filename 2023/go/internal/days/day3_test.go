package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay3Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day3-example.txt"), 4361},
		{"Actual", util.ReadFileLines("../../inputs/day3-actual.txt"), 530849},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day3Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay3Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day3-example.txt"), 467835},
		{"Actual", util.ReadFileLines("../../inputs/day3-actual.txt"), 84900879},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day3Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
