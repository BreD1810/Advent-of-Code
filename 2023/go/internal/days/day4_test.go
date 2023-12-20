package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay4Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  float64
	}{
		{"Example", util.ReadFileLines("../../inputs/day4-example.txt"), 13},
		{"Actual", util.ReadFileLines("../../inputs/day4-actual.txt"), 19135},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day4Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay4Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day4-example.txt"), 30},
		{"Actual", util.ReadFileLines("../../inputs/day4-actual.txt"), 5704953},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day4Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
