package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay5Part1(t *testing.T) {
	exampleInp := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  string
	}{
		{"Example", exampleInp, "CMZ"},
		{"Actual", util.ReadFileLines("../inputs/day05-1.txt"), "QMBMJDFTD"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day5Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %q, got %q", tc.exp, res)
			}
		})
	}
}

func TestDay5Part2(t *testing.T) {
	exampleInp := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  string
	}{
		{"Example", exampleInp, "MCD"},
		{"Actual", util.ReadFileLines("../inputs/day05-1.txt"), "NBTVTJNFJ"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day5Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %q, got %q", tc.exp, res)
			}
		})
	}
}
