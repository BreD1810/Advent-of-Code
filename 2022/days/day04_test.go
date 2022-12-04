package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay4Part1(t *testing.T) {
	exampleInp := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 2},
		{"Actual", util.ReadFileLines("../inputs/day04-1.txt"), 496},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day4Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay4Part2(t *testing.T) {
	exampleInp := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 4},
		{"Actual", util.ReadFileLines("../inputs/day04-1.txt"), 847},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day4Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
