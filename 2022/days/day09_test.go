package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay9Part1(t *testing.T) {
	exampleInp := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 13},
		{"Actual", util.ReadFileLines("../inputs/day09-1.txt"), 5735},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day9Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay9Part2(t *testing.T) {
	exampleInp := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}

	exampleInp2 := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}

	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example1", exampleInp, 1},
		{"Example2", exampleInp2, 36},
		{"Actual", util.ReadFileLines("../inputs/day09-1.txt"), 2478},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day9Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}
