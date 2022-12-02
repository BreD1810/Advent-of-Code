package days

import (
	"aoc-2022/util"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	exampleInp := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 24000},
		{"Input", util.ReadFileLines("../inputs/day01-1.txt"), 72478},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day1Part1(tc.inp)

			if res != tc.exp {
				t.Fatalf("Expected %d, got %d", tc.exp, res)
			}
		})
	}
}

func TestDay1Part2(t *testing.T) {
	exampleInp := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	tcs := []struct {
		name string
		inp  []string
		exp  int
	}{
		{"Example", exampleInp, 45000},
		{"Input", util.ReadFileLines("../inputs/day01-1.txt"), 210367},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res := day1Part2(tc.inp)

			if res != tc.exp {
				t.Fatalf("Expected %d, got %d", tc.exp, res)
			}
		})
	}
}
