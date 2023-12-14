package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay13Part1(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day13-example.txt"), 405},
		{"Actual", util.ReadFileLines("../../inputs/day13-actual.txt"), 37113},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day13Part1(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %v, got %v", tc.exp, res)
			}
		})
	}
}

func TestDay13Part2(t *testing.T) {
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", util.ReadFileLines("../../inputs/day13-example.txt"), 400},
		{"Actual", util.ReadFileLines("../../inputs/day13-actual.txt"), 30449},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day13Part2(tc.inp)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
