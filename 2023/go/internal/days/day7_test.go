package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay7Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day7-example.txt"), 6440},
		{"Actual", util.ReadFileLines("../../inputs/day7-actual.txt"), 248113761},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day7Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay7Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day7-example.txt"), 5905},
		{"Actual", util.ReadFileLines("../../inputs/day7-actual.txt"), 246285222},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day7Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
