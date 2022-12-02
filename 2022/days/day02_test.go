package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", []string{"A Y", "B X", "C Z"}, 15},
		{"Actual", util.ReadFileLines("../inputs/day02-1.txt"), 10816},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day2Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay2Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", []string{"A Y", "B X", "C Z"}, 12},
		{"Actual", util.ReadFileLines("../inputs/day02-1.txt"), 11657},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day2Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
