package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay17Part1(t *testing.T) {
	exampleInp := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>\n"

	tcs := []struct {
		name string
		inp  string
		exp  int
	}{
		{"Example", exampleInp, 3068},
		{"Actual", util.ReadFileLine("../inputs/day17-1.txt"), 3168},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day17Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay17Part2(t *testing.T) {
	exampleInp := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>\n"

	tcs := []struct {
		name string
		inp  string
		exp  int
	}{
		{"Example", exampleInp, 0},
		{"Actual", util.ReadFileLine("../inputs/day17-1.txt"), 0},
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
