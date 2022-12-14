package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay14Part1(t *testing.T) {
	exampleInp := []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 24},
		{"Actual", util.ReadFileLines("../inputs/day14-1.txt"), 719},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day14Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay14Part2(t *testing.T) {
	exampleInp := []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 93},
		{"Actual", util.ReadFileLines("../inputs/day14-1.txt"), 23390},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day14Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
