package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay19Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day19-example.txt"), 19114},
		{"Actual", util.ReadFileLines("../../inputs/day19-actual.txt"), 353046},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day19Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay19Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day19-example.txt"), 167409079868000},
		{"Actual", util.ReadFileLines("../../inputs/day19-actual.txt"), 125355665599537},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day19Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
