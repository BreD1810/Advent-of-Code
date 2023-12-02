package days

import (
	"aoc-2023/internal/util"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	exampleInp := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	tcs := []struct {
		name   string
		inp    []string
		colors cubeCounts
		exp    int
	}{
		{"Example", exampleInp, cubeCounts{red: 12, green: 13, blue: 14}, 8},
		{"Actual", util.ReadFileLines("../../inputs/day2.txt"), cubeCounts{red: 12, green: 13, blue: 14}, 2879},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day2Part1(tc.inp, tc.colors)
			if res != tc.exp {
				t.Fatalf("expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay2Part2(t *testing.T) {
	exampleInp := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	tcs := []struct {
		name   string
		inp    []string
		colors cubeCounts
		exp    int
	}{
		{"Example", exampleInp, cubeCounts{red: 12, green: 13, blue: 14}, 2286},
		{"Actual", util.ReadFileLines("../../inputs/day2.txt"), cubeCounts{red: 12, green: 13, blue: 14}, 65122},
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
